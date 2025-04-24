package models

type Employee struct {
	Id     int    `db:"id" json:"id,omitempty"`
	Name   string `db:"name" json:"name"`
	Salary int    `db:"salary" json:"salary"`
	City   string `db:"city" json:"city,omitempty"`
}
