package models

type Student struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}
