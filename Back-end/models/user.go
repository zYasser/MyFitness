package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID       uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string     `json:"name"`
	Email    string     `gorm:"unique" json:"email"`
	Username *string    `gorm:"unique" json:"username"`
	Birthday *time.Time `json:"birthday"`
	Password string     `json:"-"`
}
