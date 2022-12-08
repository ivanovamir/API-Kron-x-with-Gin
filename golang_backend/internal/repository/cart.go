package repository

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"kron-x/internal/dto"
	models2 "kron-x/internal/models"
	"kron-x/pkg"
	"strconv"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (r *CartRepository) GetUserCart(userUUID int) ([]*dto.CartProduct, error) {
	var cartProducts []*models2.CartProducts
	var cartProductsDto []*dto.CartProduct
	var cartTitle *models2.Category

	r.db.Preload(clause.Associations).
		Joins("inner join cartm2ms ug on ug.cart_products_id = cart_products.id ").
		Joins("inner join carts g on g.id= ug.cart_id ").
		Where("g.in_order = false AND g.user_id = ?", uint(userUUID)).Find(&cartProducts)

	if len(cartProducts) == 0 {
		return []*dto.CartProduct{}, nil
	} else {
		for x := range cartProducts {
			r.db.Where("Id = ?", cartProducts[x].Product.Category).Select("Title").First(&cartTitle)

			result := &dto.CartProduct{
				ID:           fmt.Sprint(cartProducts[x].Product.Id),
				Title:        cartProducts[x].Product.Title,
				VendorCode:   cartProducts[x].Product.VendorCode,
				Category:     fmt.Sprint(cartProducts[x].Product.Category),
				CategoryName: cartTitle.Title,
				Price:        fmt.Sprint(cartProducts[x].Product.Price),
				Image128:     cartProducts[x].Product.Image128,
				Count:        fmt.Sprint(cartProducts[x].Count),
			}
			cartProductsDto = append(cartProductsDto, result)
		}

		return cartProductsDto, nil
	}
}

func (r *CartRepository) UpdateUserCart(userUUID int, product *dto.UpdateCart) error {
	var cartProducts *models2.CartProducts
	var cart []models2.Cart
	var products *models2.Products
	var newTotalPrice float64

	newProductCount, _ := strconv.Atoi(product.Count)
	newProductId, _ := strconv.Atoi(product.ProductId)

	r.db.Where("id = ?", newProductId).Select("price").First(&products) //TODO bad request
	newTotalPrice = float64(newProductCount) * products.Price
	if r.db.Preload(clause.Associations).
		Joins("inner join cartm2ms ug on ug.cart_products_id = cart_products.id ").
		Joins("inner join carts g on g.id= ug.cart_id ").
		Where("g.in_order = false AND g.user_id = ?", uint(userUUID)).
		Where("product_id = ?", newProductId).First(&cartProducts).
		RowsAffected != 0 {

		cartProducts.Count = newProductCount
		cartProducts.TotalPrice = pkg.Round(newTotalPrice)
		r.db.Save(&cartProducts)
	} else {
		r.db.Where("user_id = ?", userUUID).Find(&cart)
		newCartProduct := models2.CartProducts{
			ProductID:  newProductId,
			Count:      newProductCount,
			Cart:       cart,
			CartID:     cart[0].Id,
			TotalPrice: pkg.Round(newTotalPrice),
		}
		r.db.Create(&newCartProduct)
	}
	return nil

}

func (r *CartRepository) DeleteCartProduct(userUUID int, productId int) error {
	var cartProducts *models2.CartProducts

	if err := r.db.Preload(clause.Associations).
		Joins("inner join cartm2ms ug on ug.cart_products_id = cart_products.id ").
		Joins("inner join carts g on g.id= ug.cart_id ").
		Where("g.in_order = false AND g.user_id = ?", uint(userUUID)).
		Where("product_id = ?", productId).First(&cartProducts).
		Error; err != nil {

		return ErrProdNotFound

	} else {
		r.db.Model(&cartProducts).Association("Cart").Clear()
		r.db.Unscoped().Delete(&cartProducts)
		return nil
	}

}
