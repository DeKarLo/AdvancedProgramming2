package models

import (
	"database/sql"
	"errors"
	"time"
)

type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Total     float64   `json:"total"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderRepository struct {
	DB *sql.DB
}

func (or *OrderRepository) InsertOrder(order *Order) (int, error) {
	var orderID int
	err := or.DB.QueryRow("INSERT INTO orders (user_id, total, created_at) VALUES ($1, $2, NOW()) RETURNING id",
		order.UserID, order.Total).Scan(&orderID)
	if err != nil {
		return 0, err
	}
	return orderID, nil
}

func (or *OrderRepository) GetOrderByID(orderID int) (*Order, error) {
	order := &Order{}
	err := or.DB.QueryRow("SELECT id, user_id, total, created_at FROM orders WHERE id = $1", orderID).
		Scan(&order.ID, &order.UserID, &order.Total, &order.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("order not found")
		}
		return nil, err
	}
	return order, nil
}

func (or *OrderRepository) InsertOrderItem(orderItem *OrderItem) error {
	_, err := or.DB.Exec("INSERT INTO order_items (order_id, item_id, quantity) VALUES ($1, $2, $3)",
		orderItem.OrderID, orderItem.ItemID, orderItem.Quantity)
	if err != nil {
		return err
	}
	return nil
}
