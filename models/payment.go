package models

type Payment struct {
	ID       int     `json:"id"`
	UserID   int     `json:"userID"`
	Amount   float64 `json:"amount"`
	Discount float64 `json:"discount"`
	Tax      float64 `json:"tax"`
}
