package models

type TransactionDetail struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"qty"`
}

type Transaction struct {
	Products []TransactionDetail `json:"products"`
}
