package models

type BasketProduct struct {
	ProductID int    `json:"productID"`
	SKU       string `json:"sku"`
	Quantity  int    `json:"quantity"`
}
