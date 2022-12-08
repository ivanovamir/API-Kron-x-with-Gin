package models

type Cart struct {
	Id           uint `gorm:"primaryKey"`
	UserID       uint
	InOrder      bool `gorm:"default:false"`
	Order        Order
	OrderID      uint
	CartProducts []CartProducts
}
