package postgres

import (
	"database/sql"

	"github.com/avkim12/L0/models"
	_ "github.com/lib/pq"
)

type PostgresDB interface {
	Open() error
	Close() error
	CreateOrder(p *models.Order) error
	GetOrder() ([]*models.Order, error)
}

var db *sql.DB

func Open() error {

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func Close() error {
	return db.Close()
}
