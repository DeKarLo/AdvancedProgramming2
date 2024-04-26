package models

import (
	"database/sql"
	"errors"
)

type Item struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type ItemRepository struct {
	DB *sql.DB
}

func (ir *ItemRepository) InsertItem(item *Item) error {
	_, err := ir.DB.Exec("INSERT INTO items (name, price, description) VALUES ($1, $2, $3)",
		item.Name, item.Price, item.Description)
	if err != nil {
		return err
	}
	return nil
}

func (ir *ItemRepository) GetItemByID(itemID int) (*Item, error) {
	item := &Item{}
	err := ir.DB.QueryRow("SELECT id, name, price, description FROM items WHERE id = $1", itemID).
		Scan(&item.ID, &item.Name, &item.Price, &item.Description)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("item not found")
		}
		return nil, err
	}
	return item, nil
}
