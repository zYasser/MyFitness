package repository

import "time"

type User struct {
	ID       string `gorm:"default:uuid_generate_v3()"`
	Name     string
	Email    string
	Birthday *time.Time
	Password string

}