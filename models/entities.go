package models

type Currency struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Amount int64  `json:"amount"`
}
