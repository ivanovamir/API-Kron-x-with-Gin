package models

import "github.com/spf13/viper"

type Products struct {
	Id            uint `gorm:"primaryKey"`
	Category      int64
	Title         string
	VendorCode    string
	Code1c        string
	Description   string
	Price         float64
	ImageOriginal string
	Image128      string
	Image432      string
	CanToView     bool `gorm:"default:true"`
}

func (p *Products) ProductImageOriginalPath() string {
	if p.ImageOriginal == "" {
		return p.ImageOriginal
	} else {
		return viper.GetString("files_paths.domain") + p.ImageOriginal
	}
}

func (p *Products) ProductImage128Path() string {
	if p.Image128 == "" {
		return p.Image128
	} else {
		return viper.GetString("files_paths.domain") + p.Image128
	}
}

func (p *Products) ProductImage432Path() string {
	if p.Image432 == "" {
		return p.Image432
	} else {
		return viper.GetString("files_paths.domain") + p.Image432
	}
}

type CartProducts struct {
	Id         uint `gorm:"primaryKey"`
	ProductID  int
	Count      int
	TotalPrice float64
	Cart       []Cart `gorm:"many2many:cartm2ms;"`
	CartID     uint
	OrderID    uint
	Product    Products
}

type FavouritesProducts struct {
	Id         uint `gorm:"primaryKey"`
	Product    Products
	ProductID  int
	Favourites []Favourites `gorm:"many2many:favouritesm2ms;"`
}
