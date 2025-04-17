package models

import "gorm.io/gorm"

type Workout struct {
	gorm.Model

	ID            uint   `gorm:"primaryKey;autoIncrement"`
	Name          string `json:"name"`
	Day           uint   `gorm:"not null"`
	RepLowerBound *uint
	RepUpperBound *uint
	Description   *string
	ExerciseID    uint
	ProgramId     uint
}
