package internal

import (
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Application struct {
	Router            *mux.Router
	ApplicationConfig *ApplicationConfig
}

type ApplicationConfig struct {
	Db    *gorm.DB
	Redis *redis.Client
}
