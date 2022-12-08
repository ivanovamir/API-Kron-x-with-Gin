package handler

import (
	"github.com/gin-gonic/gin"
	"kron-x/internal/dto"
	"net/http"
)

// @Summary Get all categories
// @Tags category
// @Description Get all categories
// @ID get_all_categories
// @Accept json
// @Produce json
// @Success 200 {object} dto.DefaultData{data=[]dto.Category}
// @Failure 404 {object} dto.ErrorResponse
// @Router /category [get]
func (h *Handler) GetAllCategories(c *gin.Context) {

	allCategoriesDTO, err := h.service.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": []string{}})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": allCategoriesDTO,
		})
	}
}

// @Summary Get category detail
// @Tags category
// @Description Get category detail
// @ID get_detail_category
// @Accept json
// @Produce json
// @Success 200 {object} dto.DefaultData{data=[]dto.Category}
// @Failure 404 {object} dto.ErrorResponse
// @Param cat_id query string false "Required category id"
// @Router /category_detail [get]
func (h *Handler) GetCategoriesById(c *gin.Context) {
	categoryId := c.Query("cat_id")
	categoryIdDTO := &dto.CategoryParams{
		ID: categoryId,
	}
	categoryDTO, err := h.service.GetCategoriesById(categoryIdDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": []string{}})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": categoryDTO,
		})
	}
}
