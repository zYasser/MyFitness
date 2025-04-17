package dto

type Program struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
