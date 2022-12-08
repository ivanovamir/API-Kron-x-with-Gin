package service

import (
	"kron-x/internal/dto"
	"kron-x/internal/repository"
	"kron-x/pkg"
)

type CartService struct {
	repo repository.Cart
}

func NewCartService(repo repository.Cart) *CartService {
	return &CartService{
		repo: repo,
	}
}

func (s *CartService) GetUserCart(userId string) ([]*dto.CartProduct, error) {

	userUUID, err := pkg.UserIdChecker(userId)

	if err != nil {
		return nil, err
	} else {
		return s.repo.GetUserCart(userUUID)
	}
}

func (s *CartService) UpdateUserCart(userId string, product *dto.UpdateCart) error {

	userUUID, err := pkg.UserIdChecker(userId)

	if err != nil {
		return err
	} else {
		return s.repo.UpdateUserCart(userUUID, product)
	}
}

func (s *CartService) DeleteProduct(userId string, productId int) error {

	userUUID, err := pkg.UserIdChecker(userId)

	if err != nil {
		return err
	} else {
		return s.repo.DeleteCartProduct(userUUID, productId)
	}
}
