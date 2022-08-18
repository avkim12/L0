package postgres

import (
	"database/sql"
	"encoding/json"
	"log"
)

type Order struct {
	UID   string          `json:"uid"`
	Model json.RawMessage `json:"model"`
}

type OrderDB struct {
	DB *sql.DB
}

func New() *OrderDB {

	sqldb, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=123 dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	db := OrderDB{
		DB: sqldb,
	}

	return &db
}

func (db *OrderDB) CreateOrder(order Order) error {
	
	res, err := db.DB.Exec("INSERT INTO posts(uid, model) VALUES($1,$2) RETURNING uid", order.UID, order.Model)
	if err != nil {
		return err
	}

	res.LastInsertId()

	return nil
}

func (db *OrderDB) GetOrder(uid string) (Order, error) {

	var order Order

	row := db.DB.QueryRow("SELECT * FROM orders WHERE uid = $1", uid)
	err := row.Scan(&order.UID, &order.Model)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (db *OrderDB) GetAll() ([]Order, error) {

	var orders []Order

	rows, err := db.DB.Query("SELECT * FROM orders")
	if err != nil {
		return orders, err
	}
	defer rows.Close()

	for rows.Next() {
		var order Order
		err = rows.Scan(&order.UID, &order.Model)
		if err != nil {
		  return orders, err
		}
		orders = append(orders, order)
	}

	err = rows.Err()
	if err != nil {
		return orders, err
	}

	return orders, nil
}