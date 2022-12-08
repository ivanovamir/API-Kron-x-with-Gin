package models

type Favourites struct {
	ID                 int64 `gorm:"primaryKey"`
	UserID             uint
	FavouritesProducts []FavouritesProducts `gorm:"many2many:favouritesm2ms;"`
}
