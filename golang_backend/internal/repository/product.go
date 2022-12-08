package repository

import (
	"fmt"
	"gorm.io/gorm"
	"kron-x/internal/dto"
	models2 "kron-x/internal/models"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetAllProductsByParams(params *dto.Params) ([]*dto.Product, error) {
	var allProducts []*models2.Products
	offset := (params.Page - 1) * params.Limit
	if err := r.db.Where("category IN (?)", params.CatId).Where("can_to_view",
		true).Select("Id", "Category", "Title", "Vendor_code",
		"Code1c", "Description", "Price", "Image_original", "Image128",
		"Image432").Limit(params.Limit).Offset(offset).Find(&allProducts).Error; err != nil {
		return nil, err
	} else {
		if len(allProducts) == 0 {
			return []*dto.Product{}, nil
		} else {
			var resAlready []*dto.Product
			for a := range allProducts {
				result := &dto.Product{
					Id:            fmt.Sprint(allProducts[a].Id),
					Category:      fmt.Sprint(allProducts[a].Category),
					Title:         allProducts[a].Title,
					VendorCode:    allProducts[a].VendorCode,
					Code1c:        allProducts[a].Code1c,
					Description:   allProducts[a].Description,
					Price:         fmt.Sprint(allProducts[a].Price),
					ImageOriginal: allProducts[a].ProductImageOriginalPath(),
					Image128:      allProducts[a].ProductImage128Path(),
					Image432:      allProducts[a].ProductImage432Path(),
				}
				resAlready = append(resAlready, result)
			}
			return resAlready, nil
		}
	}
}

func (r *ProductRepository) GetProductDetail(id string) (*dto.ProductInformation, error) {

	var productInfo *models2.Products
	var category *models2.Category

	if err := r.db.Where("id=?", id).Select("Id", "Category",
		"Title", "Vendor_code", "Code1c", "Description", "Price", "Image_original",
		"Image128", "Image432").First(&productInfo).Error; err != nil {
		return nil, err
	} else {

		r.db.Where("id = ?", productInfo.Category).First(&category)

		productInfoDto := &dto.ProductInformation{
			Id:            fmt.Sprint(productInfo.Id),
			Category:      fmt.Sprint(productInfo.Category),
			CategoryName:  category.Title,
			Title:         productInfo.Title,
			VendorCode:    productInfo.VendorCode,
			Code1c:        productInfo.Code1c,
			Quantity:      "1",
			Description:   productInfo.Description,
			Price:         fmt.Sprint(productInfo.Price),
			ImageOriginal: productInfo.ProductImageOriginalPath(),
			Image128:      productInfo.ProductImage128Path(),
			Image432:      productInfo.ProductImage432Path(),
		}
		return productInfoDto, nil
	}

}

func (r *ProductRepository) ParseUrlParams() []int {
	var allCategoriesIds []models2.Category

	var categoriesIds []int
	r.db.Model(&allCategoriesIds).Pluck("Id", &categoriesIds)
	return categoriesIds
}

func (r *ProductRepository) Search(params *dto.Params) ([]*dto.Product, error) {
	return nil, nil
}
