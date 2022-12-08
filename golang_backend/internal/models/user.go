package models

import (
	"time"
)

type User struct {
	Id           uint `gorm:"primaryKey"`
	CreatedAt    time.Time
	Inn          string
	Email        string
	PasswordHash string
	Phone        string
	Favourites   []Favourites
	Cart         []Cart
}

type Code struct {
	Id    uint `gorm:"primaryKey"`
	Phone string
	Email string
	Code  string
}
