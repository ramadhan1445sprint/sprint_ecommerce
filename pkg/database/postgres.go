package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/ramadhan1445sprint/sprint_ecommerce/config"
)

func NewDatabase() (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.GetString("DB_HOST"),
		config.GetString("DB_PORT"),
		config.GetString("DB_USERNAME"),
		config.GetString("DB_PASSWORD"),
		config.GetString("DB_NAME"),
		config.GetString("DB_PORT"),
	)

	db, err := sqlx.Connect("pgx", dsn)

	return db, err
}
