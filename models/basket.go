package models

type Basket struct {
	UserID   int             `json:"userID"`
	Products []BasketProduct `json:"products"`
}
