package models

type Category struct {
	Id       int64 `gorm:"primaryKey"`
	Title    string
	Image    string
	Products []Products `gorm:"foreignKey:Category"`
}
