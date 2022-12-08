package repository

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"kron-x/internal/dto"
	models2 "kron-x/internal/models"
	"strconv"
)

var (
	ErrProdAlreadyInList = fmt.Errorf("product already in list")
	ErrProdNotFound      = fmt.Errorf("product not found")
)

type FavouriteRepository struct {
	db *gorm.DB
}

func NewFavouriteRepository(db *gorm.DB) *FavouriteRepository {
	return &FavouriteRepository{
		db: db,
	}
}

func (r *FavouriteRepository) GetUserFavourites(userUUID int) ([]*dto.FavouritesProducts, error) {
	var favouriteProducts []*models2.FavouritesProducts
	var favouriteProductsDto []*dto.FavouritesProducts
	var cartTitle *models2.Category

	r.db.Preload(clause.Associations).
		Joins("inner join favouritesm2ms ug on ug.favourites_products_id = favourites_products.id ").
		Joins("inner join favourites g on g.id= ug.favourites_id ").
		Where("g.user_id = ?", uint(userUUID)).Find(&favouriteProducts)

	if len(favouriteProducts) == 0 {
		return []*dto.FavouritesProducts{}, nil
	} else {
		for x := range favouriteProducts {
			r.db.Where("Id = ?", favouriteProducts[x].Product.Category).Select("Title").First(&cartTitle)
			result := &dto.FavouritesProducts{
				ID:           fmt.Sprint(favouriteProducts[x].Product.Id),
				Title:        favouriteProducts[x].Product.Title,
				VendorCode:   favouriteProducts[x].Product.VendorCode,
				Category:     fmt.Sprint(favouriteProducts[x].Product.Category),
				CategoryName: cartTitle.Title,
				Price:        fmt.Sprint(favouriteProducts[x].Product.Price),
				Image128:     favouriteProducts[x].Product.Image128,
			}
			favouriteProductsDto = append(favouriteProductsDto, result)

		}
		return favouriteProductsDto, nil
	}
}

func (r *FavouriteRepository) UpdateUserFavourites(userUUID int, product *dto.UpdateFavourite) error {
	var favouriteProducts *models2.FavouritesProducts
	var favourite []models2.Favourites

	newProductId, _ := strconv.Atoi(product.ProductId)

	if r.db.Preload(clause.Associations).
		Joins("inner join favouritesm2ms ug on ug.favourites_products_id = favourites_products.id ").
		Joins("inner join favourites g on g.id= ug.favourites_id ").
		Where("g.user_id = ?", uint(userUUID)).
		Where("product_id = ?", newProductId).Find(&favouriteProducts).
		RowsAffected != 0 {

		return ErrProdAlreadyInList
	} else {
		r.db.Where("user_id = ?", userUUID).Find(&favourite)
		newCartProduct := models2.FavouritesProducts{
			ProductID:  newProductId,
			Favourites: favourite,
		}
		r.db.Create(&newCartProduct)
		return nil
	}

}

func (r *FavouriteRepository) DeleteFavouriteProduct(userUUID int, productId int) error {
	var favouritesProducts *models2.FavouritesProducts

	if err := r.db.Preload(clause.Associations).
		Joins("inner join favouritesm2ms ug on ug.favourites_products_id = favourites_products.id ").
		Joins("inner join favourites g on g.id= ug.favourites_id ").
		Where("g.user_id = ?", uint(userUUID)).
		Where("product_id = ?", productId).First(&favouritesProducts).
		Error; err != nil {

		return ErrProdNotFound

	} else {
		r.db.Model(&favouritesProducts).Association("Favourites").Clear()
		r.db.Unscoped().Delete(&favouritesProducts)
		return nil
	}

}
