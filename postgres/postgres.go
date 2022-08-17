package postgres

import (
	"database/sql"
	"encoding/json"
)

type Order struct {
	UID   string          `json:"uid"`
	Model json.RawMessage `json:"model"`
}

type OrderDB struct {
	DB *sql.DB
}

func (db OrderDB) CreateOrder(order Order) error {
	res, err := db.DB.Exec(insertOrderSchema, order.UID, order.Model)
	if err != nil {
		return err
	}
	res.LastInsertId()
	return nil
}

func (db OrderDB) GetOrder(uid string) (Order, error) {
	var order Order
	row := db.DB.QueryRow(selectOrderSchema, uid)
	err := row.Scan(&order.UID, &order.Model)
	if err != nil {
		return order, err
	}
	return order, nil
}
