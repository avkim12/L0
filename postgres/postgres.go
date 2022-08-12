package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/avkim12/L0/cache"
	"github.com/avkim12/L0/model"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "postgres"
)

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func CreateOrder(model model.Order) (sql.Result, error) {
	
	db := OpenConnection()
	
	sqlStatement := `INSERT INTO users (id, model) VALUES ($1, $2)`
	res, err := db.Exec(sqlStatement, model.Id, model.Model)
	if err != nil {
		return nil, err
	}
	cache := cache.New(5 * time.Minute, 10 * time.Minute)
	cache.Set(model.Id, model.Model, 5 * time.Minute)

	return res, nil
}