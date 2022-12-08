package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kron-x/internal/dto"
	"net/http"
	"strconv"
)

// @Security ApiKeyAuth
// @Summary Get user cart
// @Tags cart
// @Description Get user cart
// @ID get_user_cart
// @Accept json
// @Produce json
// @Success 200 {object} dto.DefaultData{data=[]dto.CartProduct}
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/cart [get]
func (h *Handler) GetUserCart(c *gin.Context) {
	id, _ := c.Get(userCtx)
	userCart, err := h.service.GetUserCart(fmt.Sprint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": []string{},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": userCart,
		})
	}
}

// @Security ApiKeyAuth
// @Summary Update user cart
// @Tags cart
// @Description Update user cart
// @ID update_user_cart
// @Accept json
// @Produce json
// @Param input body dto.UpdateCart true "new product info"
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/cart/add [post]
func (h *Handler) UpdateUserCart(c *gin.Context) {
	id, _ := c.Get(userCtx)
	var product *dto.UpdateCart

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong json input",
		})
	} else {
		err := h.service.UpdateUserCart(fmt.Sprint(id), product)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprint(err),
			})
		}
	}
}

// @Security ApiKeyAuth
// @Summary Delete item from user cart
// @Tags cart
// @Description Delete item from user cart
// @ID delete_user_cart
// @Accept json
// @Produce json
// @Param input body dto.DeleteCart true "delete product info"
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/cart/del [post]
func (h *Handler) DeleteProduct(c *gin.Context) {
	id, _ := c.Get(userCtx)

	var product *dto.DeleteCart

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong json input",
		})
	} else {
		productId, _ := strconv.Atoi(product.ProductId)
		err := h.service.DeleteProduct(fmt.Sprint(id), productId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprint(err),
			})
		}
	}
}
