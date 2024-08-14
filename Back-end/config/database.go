package config

import (
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"

	"github.com/zYasser/MyFitness/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	DATABASE_USERNAME := utils.GetEnv("DATABASE_USER")
	DATABASE_PASSWORD := utils.GetEnv("DATABASE_PASSWORD")
	fmt.Printf(DATABASE_PASSWORD, "\n ", DATABASE_USERNAME)
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=MyFitness port=5432 sslmode=disable", DATABASE_USERNAME, DATABASE_PASSWORD)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Connection Failed : %v", err)
	}
	log.Println("Connection has been established")

	return db

}


func InitRedis() *redis.Client{	

	
	redis := redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "", // no password set
		DB:		  0,  // use default DB
	})
	return redis
}