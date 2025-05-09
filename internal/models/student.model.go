package models

import "mime/multipart"

type Student struct {
	Id       int    `db:"id" json:"id,omitempty"`
	Name     string `db:"name" json:"username"`
	Password string `db:"password,omitempty" json:"password"`
	Role     string `db:"role,omitempty" json:"role,omitempty"`
	Image    string `db:"image" json:"image"`
}

type StudentForm struct {
	Name  string                `form:"username"`
	Image *multipart.FileHeader `form:"img"`
}
