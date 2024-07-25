package dto

import (
	"time"

)
type User struct{
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required"`
	Birthday *time.Time `json:"birthday"`
	Password string `json:"password"`
}
