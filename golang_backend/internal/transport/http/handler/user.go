package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"       // swagger embed files
	_ "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"kron-x/internal/dto"
	"net/http"
)

// @Summary Sign-up
// @Tags auth
// @Description Register user
// @ID create-account
// @Accept json
// @Produce json
// @Param input body dto.RegisterUser true "account info"
// @Success 200 {object} dto.Token
// @Failure 401 {object} dto.ErrorResponse
// @Router /auth/sign_up [post]
func (h *Handler) RegisterUser(c *gin.Context) {
	var userDTO *dto.RegisterUser

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "wrong json input",
		})
	} else {
		token, err := h.service.RegisterUser(userDTO)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": fmt.Sprint(err),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		}
	}
}

// @Summary Sign-in
// @Tags auth
// @Description Authenticate user
// @ID authenticate-account
// @Accept json
// @Produce json
// @Param input body dto.UserAuth true "account info"
// @Success 200 {object} dto.Token
// @Failure 401 {object} dto.ErrorResponse
// @Router /auth/sign_in [post]
func (h *Handler) AuthenticateUser(c *gin.Context) {
	var userDTO *dto.UserAuth

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "wrong json input",
		})
	} else {
		token, err := h.service.AuthenticateUser(userDTO)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": fmt.Sprint(err),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		}
	}
}

// @Summary Change password
// @Tags auth
// @Description Change user password
// @ID new_password-account
// @Accept json
// @Produce json
// @Param input body dto.UserInformation true "account info"
// @Success 200
// @Failure 401 {object} dto.ErrorResponse
// @Router /auth/new_password [post]
func (h *Handler) ChangeUserPassword(c *gin.Context) {
	var userInfoDto *dto.UserInformation

	if err := c.ShouldBindJSON(&userInfoDto); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "wrong json input",
		})
	} else {
		err := h.service.ChangeUserPassword(userInfoDto)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": fmt.Sprint(err),
			})
		}
	}
}

// @Summary Generate sms-code
// @Tags auth
// @Description Send user code for registration or authentication
// @ID generate_phone_code
// @Accept json
// @Produce json
// @Param input body dto.CodeGenerator true "user info"
// @Success 200
// @Failure 401 {object} dto.ErrorResponse
// @Router /auth/code [post]
func (h *Handler) CodeGenerator(c *gin.Context) {
	var code *dto.CodeGenerator
	if err := c.ShouldBindJSON(&code); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "wrong json input",
		})
	} else {
		if err := h.service.CodeGenerator(code); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": fmt.Sprint(err),
			})
		}
	}
}
