package service

import (
	"kron-x/internal/dto"
	"kron-x/internal/repository"
)

type CatalogService struct {
	repo repository.Catalog
}

func NewCatalogService(repo repository.Catalog) *CatalogService {
	return &CatalogService{
		repo: repo,
	}
}

func (s *CatalogService) GetAllCatalogs(params *dto.Params) ([]*dto.Catalog, error) {
	return s.repo.GetAllCatalogs(params)
}
