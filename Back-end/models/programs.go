package models

import "gorm.io/gorm"

type Program struct {
	gorm.Model

	ID          int    `gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name"`
	Description *string
	Workout     []Workout
}
