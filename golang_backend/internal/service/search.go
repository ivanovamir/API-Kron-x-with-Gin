package service

import (
	"kron-x/internal/dto"
	"kron-x/internal/repository"
)

type SearchService struct {
	repo repository.Search
}

func NewSearchService(repo repository.Search) *SearchService {
	return &SearchService{
		repo: repo,
	}
}

func (s *SearchService) GetSearch(searchParams *dto.Params) ([]*dto.Product, error) {
	//TODO implement morphological logic
	return s.repo.GetSearch(searchParams)
}
