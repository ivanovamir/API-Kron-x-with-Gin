package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kron-x/internal/dto"
	"net/http"
	"strconv"
)

// @Security ApiKeyAuth
// @Summary Get user favourites list
// @Tags favourite
// @Description Get user favourites list
// @ID get_user_favourite
// @Accept json
// @Produce json
// @Success 200 {object} dto.DefaultData{data=[]dto.FavouritesProducts}
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/favourite [get]
func (h *Handler) GetUserFavourites(c *gin.Context) {
	id, _ := c.Get(userCtx)
	userFavourites, err := h.service.GetUserFavourites(fmt.Sprint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": []string{},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": userFavourites,
		})
	}
}

// @Security ApiKeyAuth
// @Summary Update user favourites list
// @Tags favourite
// @Description Update user favourites list
// @ID update_user_favourite
// @Accept json
// @Produce json
// @Param input body dto.UpdateFavourite true "new favourite product info"
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/favourite/add [post]
func (h *Handler) UpdateUserFavourites(c *gin.Context) {
	id, _ := c.Get(userCtx)

	var product *dto.UpdateFavourite

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong json input",
		})
	} else {
		err := h.service.UpdateUserFavourites(fmt.Sprint(id), product)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprint(err),
			})
		}
	}
}

// @Security ApiKeyAuth
// @Summary Delete item from user favourites list
// @Tags favourite
// @Description Delete item from user favourites list
// @ID delete_user_favourite
// @Accept json
// @Produce json
// @Param input body dto.DeleteFavourite true "delete favourite product info"
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/favourite/del [post]
func (h *Handler) DeleteFavouriteProduct(c *gin.Context) {
	id, _ := c.Get(userCtx)

	var product *dto.DeleteFavourite

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong json input",
		})
	} else {
		productId, _ := strconv.Atoi(product.ProductId)
		err := h.service.DeleteFavouriteProduct(fmt.Sprint(id), productId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprint(err),
			})
		}
	}
}
