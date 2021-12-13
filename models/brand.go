package models

type Brand struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	ProductQty int     `json:"productQty"`
	TotalWorth float64 `json:"totalWorth"`
}
