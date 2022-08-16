package postgres

import "github.com/avkim12/L0/models"

func CreateOrder(order models.Order) error {
	res, err := db.Exec(insertOrderSchema, order.Id, order.Model)
	if err != nil {
		return err
	}
	res.LastInsertId()
	return nil
}

func GetOrder(uid string) (models.Order, error) {
	var order models.Order
	row := db.QueryRow(selectOrderSchema, uid)
	err := row.Scan(&order.Id, &order.Model)
	if err != nil {
		return order, err
	}
	return order, nil
}