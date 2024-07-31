package service

import "gorm.io/gorm"

type Exercise struct {
	gorm.Model

	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `json:"name" gorm:"not null"`
	Type string `json:"type" gorm:"not null"`
}



