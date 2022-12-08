package repository

import (
	"fmt"
	"gorm.io/gorm"
	"kron-x/internal/dto"
	"kron-x/internal/models"
	"kron-x/pkg"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) GetAllCategories() ([]*dto.Category, error) {
	var allCategories []*models.Category
	if err := r.db.Find(&allCategories).Error; err != nil {
		return nil, err
	} else {
		if len(allCategories) == 0 {
			return []*dto.Category{}, nil
		} else {
			var resAlready []*dto.Category
			for a := range allCategories {
				result := &dto.Category{
					ID:         fmt.Sprint(allCategories[a].Id),
					Title:      allCategories[a].Title,
					Image:      pkg.PhotoChecker(allCategories[a].Image),
					CountItems: "0",
				}
				resAlready = append(resAlready, result)
			}
			return resAlready, nil
		}
	}
}

func (r *CategoryRepository) GetCategoriesById(id *dto.CategoryParams) ([]*dto.Category, error) {
	var category *models.Category
	if err := r.db.Where("id=?", id.ID).First(&category).Error; err != nil {
		return nil, err
	} else {
		var resAlready []*dto.Category
		result := &dto.Category{
			ID:         fmt.Sprint(category.Id),
			Title:      category.Title,
			Image:      pkg.PhotoChecker(category.Image),
			CountItems: "0",
		}
		resAlready = append(resAlready, result)
		return resAlready, nil
	}
}
