package models

import "time"

type Order struct {
	Id                         uint `gorm:"primaryKey"`
	Note                       string
	OptionOfDeliveryAndPayment *OptionOfDeliveryAndPayment
	UserID                     uint
	OrderStatusID              uint `gorm:"default:1"`
	CartID                     uint
	CreatedAt                  time.Time
}

type OrderStatus struct {
	Id    uint `gorm:"primaryKey"`
	Name  string
	Order Order
}

type OptionOfDeliveryAndPayment struct {
	Id              uint `gorm:"primaryKey"`
	City            string
	Street          string
	House           string
	Frame           string
	Entrance        string
	ApartmentOffice string
	OrderID         uint
}
