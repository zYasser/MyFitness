package models

import "gorm.io/gorm"

type Exercise struct {
	gorm.Model

	Name    string `json:"name" gorm:"not null"`
	Type    string `json:"type" gorm:"not null"`
	Workout []Workout
}
