package models

type Customer struct {
	ID             int     `json:"id"`
	UserID         int     `json:"userID"`
	OrderQty       int     `json:"orderQty"`
	PurchaseAmount float64 `json:"purchaseAmount"`
}
