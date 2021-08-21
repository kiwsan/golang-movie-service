package models

type Movie struct {
	Id       int64  `json:"id"`
	Name     string `json:"name" validate:"required,min=2,max=50"`
	Director string `json:"director" validate:"required,min=2,max=50"`
	Writer   string `json:"writer" validate:"required,min=2,max=50"`
	Stars    string `json:"stars" validate:"required,min=2,max=50"`
}
