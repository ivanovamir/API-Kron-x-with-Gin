package repository

import (
	"fmt"
	"gorm.io/gorm"
	"kron-x/internal/dto"
	"kron-x/internal/models"
)

type CatalogRepository struct {
	db *gorm.DB
}

func NewCatalogRepository(db *gorm.DB) *CatalogRepository {
	return &CatalogRepository{
		db: db,
	}
}

func (r *CatalogRepository) GetAllCatalogs(params *dto.Params) ([]*dto.Catalog, error) {
	var catalogs []*models.Catalog
	offset := (params.Page - 1) * params.Limit
	if err := r.db.Find(&catalogs).Limit(params.Limit).Offset(offset).Error; err != nil {
		return nil, err
	} else {
		if len(catalogs) == 0 {
			return []*dto.Catalog{}, nil
		} else {
			var catalogsDto []*dto.Catalog
			for a := range catalogs {
				result := &dto.Catalog{
					Id:    fmt.Sprint(catalogs[a].Id),
					Title: catalogs[a].Title,
					File:  catalogs[a].CatalogFilePath(),
				}
				catalogsDto = append(catalogsDto, result)
			}
			return catalogsDto, nil
		}
	}
}
