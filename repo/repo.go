package repo

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/jmoiron/sqlx"
)

type RepoInterface interface {
	CreateProduct(product Product) error
	GetDetailProduct(id uuid.UUID) (Product, error)
	UpdateProduct(product Product) error
	DeleteProduct(id uuid.UUID) error
	GetListProduct(keys Key) ([]Product, error)
}

func NewRepo(db *sqlx.DB) RepoInterface {
	return &repo{db: db}
}

type repo struct {
	db *sqlx.DB
}

type Product struct {
	ID            uuid.UUID `db:"id" json:"id" validate:"required,uuid4"`
	Name          string    `db:"name" json:"name" validate:"required,max=60,min=5"`
	Price         float64   `db:"price" json:"price" validate:"required,numeric,gte=0"`
	Stock         int       `db:"stock" json:"stock" validate:"required,numeric,gte=0"`
	ImageUrl      string    `db:"image_url" json:"imageUrl" validate:"required"`
	Condition     string    `db:"condition" json:"condition" validate:"required,validCondition"`
	IsPurchasable bool      `db:"is_purchasable" json:"isPurchasable"`
	Tags          []string  `db:"tags" json:"tags" validate:"required"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}

type Key struct {
	UserOnly       *bool    `json:"userOnly"`
	Limit          *int     `json:"limit"`
	Offset         *int     `json:"offset"`
	Tags           []string `json:"tags"`
	Condition      *string  `json:"condition"`
	ShowEmptyStock *bool    `json:"showEmptyStock"`
	MaxPrice       *float64 `json:"maxPrice"`
	MinPrice       *float64 `json:"minPrice"`
	SortBy         *string  `json:"sortBy"`
	OrderBy        *string  `json:"orderBy"`
	Search         *string  `json:"search"`
}

func (r *repo) CreateProduct(product Product) error {
	query := `INSERT INTO product (id, name, price, stock, image_url, condition, is_purchasable, tags, created_at, updated_at)
               VALUES (:id, :name, :price, :stock, :image_url, :condition, :is_purchasable, :tags, :created_at, :updated_at)`

	_, err := r.db.NamedExec(query, &product)
	if err != nil {
		log.Println("Error executing query:", err)
		return err
	}

	return nil
}

func (r *repo) GetDetailProduct(id uuid.UUID) (Product, error) {
	var product Product
	var tags pgtype.VarcharArray
	query := "SELECT tags FROM product WHERE id = $1"

	// Query for a single row
	err := r.db.QueryRowx(query, id).Scan(&tags)
	if err != nil {
		log.Println("Error executing query:", err)
		return product, err
	}

	// Extract tags from VarcharArray
	var tagsSlice []string
	for _, tag := range tags.Elements {
		if tag.Status != pgtype.Null {
			tagsSlice = append(tagsSlice, string(tag.String))
		}
	}

	query1 := "SELECT id, name, price, stock, image_url, condition, is_purchasable FROM product WHERE id = $1"
	err = r.db.QueryRowx(query1, id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Stock,
		&product.ImageUrl,
		&product.Condition,
		&product.IsPurchasable,
	)
	if err != nil {
		log.Println("Error executing query:", err)
		return product, err
	}

	product.Tags = tagsSlice

	return product, nil
}

func (r *repo) UpdateProduct(product Product) error {
	query := `UPDATE product set name = :name, price = :price, stock = :stock, image_url = :image_url,
							condition = :condition, is_purchasable = :is_purchasable, tags = :tags, updated_at = :updated_at
							WHERE id = :id`

	_, err := r.db.NamedExec(query, &product)
	if err != nil {
		log.Println("Error executing query:", err)
		return err
	}

	return nil
}

func (r *repo) DeleteProduct(id uuid.UUID) error {
	query := `DELETE FROM product WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("Error executing query:", err)
		return err
	}

	return nil
}

func (r *repo) GetListProduct(keys Key) ([]Product, error) {
	// var tags pgtype.VarcharArray
	var conditions []string

	if keys.MaxPrice != nil && keys.MinPrice != nil {
		conditions = append(conditions, fmt.Sprintf("price BETWEEN %.2f AND %.2f", *keys.MinPrice, *keys.MaxPrice))
	}

	if keys.Condition != nil {
		conditions = append(conditions, fmt.Sprintf("condition = '%s'", *keys.Condition))
	}

	if keys.ShowEmptyStock != nil {
		if *keys.ShowEmptyStock {
			conditions = append(conditions, "stock = 0")
		} else {
			conditions = append(conditions, "stock > 0")
		}
	}

	if len(keys.Tags) > 0 {
		var tagConditions []string
		for _, tag := range keys.Tags {
			tagConditions = append(tagConditions, fmt.Sprintf("'%s'", tag))
		}
		conditions = append(conditions, fmt.Sprintf("ARRAY[%s] && tags", strings.Join(tagConditions, ",")))
	}

	if keys.Search != nil {
		conditions = append(conditions, fmt.Sprintf("name LIKE '%%%s%%'", *keys.Search))
	}

	// Check if any conditions were provided
	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	orderByClause := ""
	if keys.OrderBy != nil && keys.SortBy != nil {
		orderByClause = fmt.Sprintf("ORDER BY %s %s", *keys.SortBy, *keys.OrderBy)
	}

	limitClause := ""
	if keys.Limit != nil {
		limitClause = fmt.Sprintf("LIMIT %d", *keys.Limit)
	}

	offsetClause := ""
	if keys.Offset != nil {
		offsetClause = fmt.Sprintf("OFFSET %d", *keys.Offset)
	}

	query := fmt.Sprintf(`
        SELECT id, name, price, stock, image_url, condition, is_purchasable
        FROM product
        %s
        %s
        %s
        %s`, whereClause, orderByClause, limitClause, offsetClause)

	fmt.Println(query)

	// Execute the query
	var products []Product
	err := r.db.Select(&products, query)
	if err != nil {
		log.Println("Error executing query:", err)
		return products, err
	}

	return products, nil
}
