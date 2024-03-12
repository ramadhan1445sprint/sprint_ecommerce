package repo

import (
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

func (r *repo) CreateProduct(product Product) error {
	query := `INSERT INTO product (id, name, price, stock, image_url, condition, is_purchasable, tags, created_at, updated_at)
               VALUES (:id, :name, :price, :stock, :image_url, :condition, :is_purchasable, :tags, :created_at, :updated_at)`

	_, err := r.db.NamedExec(query, &product)
	if err != nil {
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
		return err
	}

	return nil
}

func (r *repo) DeleteProduct(id uuid.UUID) error {
	query := `DELETE FROM product WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
