package service

import (
	"kron-x/internal/dto"
	"kron-x/internal/repository"
	"kron-x/pkg"
)

type FavouriteService struct {
	repo repository.Favourite
}

func NewFavouriteService(repo repository.Favourite) *FavouriteService {
	return &FavouriteService{
		repo: repo,
	}
}

func (s *FavouriteService) GetUserFavourites(userId string) ([]*dto.FavouritesProducts, error) {

	userUUID, err := pkg.UserIdChecker(userId)

	if err != nil {
		return nil, err
	} else {
		return s.repo.GetUserFavourites(userUUID)
	}
}

func (s *FavouriteService) UpdateUserFavourites(userId string, product *dto.UpdateFavourite) error {

	userUUID, err := pkg.UserIdChecker(userId)

	if err != nil {
		return err
	} else {
		return s.repo.UpdateUserFavourites(userUUID, product)
	}
}

func (s *FavouriteService) DeleteFavouriteProduct(userId string, productId int) error {
	userUUID, err := pkg.UserIdChecker(userId)

	if err != nil {
		return err
	} else {
		return s.repo.DeleteFavouriteProduct(userUUID, productId)
	}
}
