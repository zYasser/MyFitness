package service

import (
	// "github.com/zYasser/MyFitness/utils"
	service2 "github.com/zYasser/MyFitness/internal/exercise/service"
	"gorm.io/gorm"
)

// var logger = utils.GetLogger()

func Migration(db *gorm.DB) {
	db.AutoMigrate(&User{}, &service2.Exercise{})
}
