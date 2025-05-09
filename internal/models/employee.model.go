package models

type Employee struct {
	Id     int    `db:"id" json:"id,omitempty"`
	Name   string `db:"name" json:"name" binding:"required"`
	Salary int    `db:"salary" json:"salary" binding:"required,gt=10"`
	City   string `db:"city" json:"city,omitempty" binding:"required"`
}
