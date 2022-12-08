package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Search
// @Tags search
// @Description Get product by search
// @ID get_product_by_search
// @Produce json
// @Success 200 {object} dto.DefaultData{data=[]dto.Product}
// @Failure 404 {object} dto.ErrorResponse
// @Param s query string false "Search text"
// @Router /search [get]
func (h *Handler) GetSearch(c *gin.Context) {
	allParams := h.ParseUrlParams(c)
	productsDto, err := h.service.GetSearch(allParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": fmt.Sprint(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": productsDto,
		})
	}
}
