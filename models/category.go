package models

type Category struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ProductQty int    `json:"productQty"`
	IsMain     bool   `json:"isMain"`
}
