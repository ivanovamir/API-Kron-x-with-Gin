package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kron-x/internal/dto"
	"net/http"
)

// @Security ApiKeyAuth
// @Summary Create order
// @Tags order
// @Description Create order
// @ID create_user_order
// @Accept json
// @Produce json
// @Param input body dto.CreateOrder true "order info"
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/order [post]
func (h *Handler) CreateOrder(c *gin.Context) {
	id, _ := c.Get(userCtx)
	var orderDto *dto.CreateOrder

	if err := c.ShouldBindJSON(&orderDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong json input",
		})
	} else {
		err := h.service.CreateOrder(fmt.Sprint(id), orderDto)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprint(err),
			})
		}
	}
}

// @Security ApiKeyAuth
// @Summary Get user order
// @Tags order
// @Description Get user order
// @ID get_user_order
// @Accept json
// @Produce json
// @Success 200 {object} dto.DefaultData{data=[]dto.GetOrder}
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/order [get]
func (h *Handler) GetUserOrder(c *gin.Context) {

	id, _ := c.Get(userCtx)

	orderDto, err := h.service.GetUserOrder(fmt.Sprint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprint(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": orderDto,
		})
	}
}
