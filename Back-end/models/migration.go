package models

import (
	// "github.com/zYasser/MyFitness/utils"
	"gorm.io/gorm"
)

// var logger = utils.GetLogger()

func Migration(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Exercise{}, &Program{}, &Workout{})
}
