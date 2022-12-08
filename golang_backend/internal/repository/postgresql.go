package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	models2 "kron-x/internal/models"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func NewPostgresDB(cfg Config) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(fmt.Sprintf(fmt.Sprintf("postgres://%s:%s@%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.DBName))), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		//Logger:                 logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	} else {
		if err := db.AutoMigrate(&models2.OrderStatus{}, &models2.Category{}, &models2.Products{}, &models2.Slider{}, &models2.About{}, &models2.Metric{},
			&models2.Requisite{}, &models2.HeadOffice{}, &models2.Feature{}, &models2.Email{}, &models2.Service{},
			&models2.User{}, &models2.Cart{}, &models2.Favourites{}, &models2.Code{}, &models2.CartProducts{}, &models2.FavouritesProducts{},
			&models2.Order{}, &models2.OptionOfDeliveryAndPayment{}, &models2.FeedbackCall{}, &models2.Catalog{},
			&models2.FeedbackSelection{}, &models2.Support{}); err != nil {
			return nil, err
		} else {
			return db, nil
		}
	}
}
