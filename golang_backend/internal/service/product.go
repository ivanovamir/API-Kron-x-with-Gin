package service

import (
	"kron-x/internal/dto"
	"kron-x/internal/repository"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) GetAllProductsByParams(params *dto.Params) ([]*dto.Product, error) {
	return s.repo.GetAllProductsByParams(params)
}

func (s *ProductService) GetProductDetail(id string) (*dto.ProductInformation, error) {
	return s.repo.GetProductDetail(id)
}

func (s *ProductService) ParseUrlParams() []int {
	return s.repo.ParseUrlParams()
}
