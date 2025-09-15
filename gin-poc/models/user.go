package models

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required,min=2"`
}
