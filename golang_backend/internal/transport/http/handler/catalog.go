package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get all catalogs
// @Tags catalog
// @Description Get all catalogs
// @ID get_all_catalogs
// @Accept json
// @Produce json
// @Success 200 {object} dto.DefaultData{data=[]dto.Catalog}
// @Failure 404 {object} dto.ErrorResponse
// @Param limit query string false "Catalog limit in response"
// @Param page query string false "Response page"
// @Router /catalog [get]
func (h *Handler) GetAllCatalogs(c *gin.Context) {
	allParams := h.ParseUrlParams(c)
	catalogDto, err := h.service.GetAllCatalogs(allParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": []string{},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": catalogDto,
		})
	}
}
