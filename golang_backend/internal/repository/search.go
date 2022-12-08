package repository

import (
	"fmt"
	"gorm.io/gorm"
	"kron-x/internal/dto"
	"kron-x/internal/models"
)

type SearchRepository struct {
	db *gorm.DB
}

func NewSearchRepository(db *gorm.DB) *SearchRepository {
	return &SearchRepository{
		db: db,
	}
}

func (r *SearchRepository) GetSearch(searchParams *dto.Params) ([]*dto.Product, error) {
	var products []*models.Products
	var productsDto []*dto.Product
	offset := (searchParams.Page - 1) * searchParams.Limit
	if err := r.db.Where("title LIKE ?", "%"+searchParams.Search+"%").Or("vendor_code LIKE ?", "%"+searchParams.Search+"%").Limit(searchParams.Limit).Offset(offset).Find(&products).
		Error; err != nil {
		return []*dto.Product{}, nil
	} else {
		if len(products) == 0 {
			return []*dto.Product{}, nil
		} else {
			for x := range products {
				productDto := &dto.Product{
					Id:            fmt.Sprint(products[x].Id),
					Category:      fmt.Sprint(products[x].Category),
					Title:         products[x].Title,
					VendorCode:    fmt.Sprint(products[x].VendorCode),
					Code1c:        fmt.Sprint(products[x].Code1c),
					Description:   products[x].Description,
					Price:         fmt.Sprint(products[x].Price),
					ImageOriginal: products[x].ProductImageOriginalPath(),
					Image128:      products[x].ProductImage128Path(),
					Image432:      products[x].ProductImage432Path(),
				}
				productsDto = append(productsDto, productDto)
			}
			return productsDto, nil
		}
	}
}
