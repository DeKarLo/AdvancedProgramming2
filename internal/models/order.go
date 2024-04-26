package models

import (
	"database/sql"
	"errors"
	"time"
)

// Order represents an order in the system
type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Total     float64   `json:"total"`
	CreatedAt time.Time `json:"created_at"`
}

// OrderRepository represents a repository for order operations
type OrderRepository struct {
	DB *sql.DB
}

// InsertOrder inserts a new order into the database
func (or *OrderRepository) InsertOrder(order *Order) error {
	_, err := or.DB.Exec("INSERT INTO orders (user_id, total, created_at) VALUES ($1, $2, $3)",
		order.UserID, order.Total, order.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

// GetOrderByID retrieves an order from the database by ID
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
