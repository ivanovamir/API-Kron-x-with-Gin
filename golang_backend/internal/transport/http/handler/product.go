package handler

import (
	"github.com/gin-gonic/gin"
	"kron-x/internal/dto"
	"net/http"
)

// @Summary Get all products
// @Tags product
// @Description Get all products with params
// @ID get_all_products
// @Accept json
// @Produce json
// @Success 200 {object} dto.DefaultData{data=[]dto.Product}
// @Failure 404 {object} dto.ErrorResponse
// @Param cat_id query string false "Products belonging to this category"
// @Param limit query string false "Product limit in response"
// @Param page query string false "Response page"
// @Router /products [get]
func (h *Handler) GetAllProductsByParams(c *gin.Context) {

	allParams := h.ParseUrlParams(c)
	allProducts, err := h.service.GetAllProductsByParams(allParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": []string{},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": allProducts,
		})
	}
}

// @Summary Get product detail
// @Tags product
// @Description Get product detail
// @ID get_detail_product
// @Accept json
// @Produce json
// @Success 200 {object} dto.DefaultData{data=[]dto.ProductInformation}
// @Failure 404 {object} dto.ErrorResponse
// @Param prod_id query string false "Required product id"
// @Router /product_detail [get]
func (h *Handler) GetProductDetail(c *gin.Context) {
	productId := c.Query("prod_id")
	productInfo, err := h.service.GetProductDetail(productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": []string{},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": []*dto.ProductInformation{productInfo},
		})
	}
}
