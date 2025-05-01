package models

type Student struct {
	Id       int    `db:"id"`
	Name     string `db:"name" json:"username"`
	Password string `db:"password,omitempty" json:"password"`
	Role     string `db:"role,omitempty" json:"role,omitempty"`
}
