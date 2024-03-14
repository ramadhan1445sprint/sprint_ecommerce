package repo

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/ramadhan1445sprint/sprint_ecommerce/entity"
)

type StockRepoInterface interface {
	UpdateStock(product *entity.Product) error
}

func NewStockRepo(db *sqlx.DB) StockRepoInterface {
	return &stockRepo{db: db}
}

type stockRepo struct {
	db *sqlx.DB
}

func (r *stockRepo) UpdateStock(product *entity.Product) error {
	res, err := r.db.Exec("UPDATE products SET stock = $1 where id = $2", product.Stock, product.ID)

	if err != nil {
		return err
	}

	rowsEffected, _ := res.RowsAffected()

	if rowsEffected == 0 {
		return errors.New("bank account not found")
	}

	return nil
}
