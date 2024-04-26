package models

type OrderItem struct {
	ID       int `json:"id"`
	OrderID  int `json:"order_id"`
	ItemID   int `json:"item_id"`
	Quantity int `json:"quantity"`
}
