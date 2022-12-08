package handler

import (
	"github.com/gin-gonic/gin"
	"kron-x/internal/dto"
	"net/http"
	"strconv"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	headerParts := strings.Split(header, " ")

	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "empty auth user"})
	} else if len(headerParts) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "empty auth user"})
	}

	userId, err := h.service.ParseToken(headerParts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "empty auth user"})
	}

	c.Set(userCtx, userId)
}

func (h *Handler) ParseUrlParams(c *gin.Context) *dto.Params {
	allCategoriesId := h.service.ParseUrlParams()

	limit := 0
	page := 1
	var search string
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		case "cat_id":
			allCategoriesId = toIntArray(queryValue)
		case "s":
			search = queryValue

		}
	}
	return &dto.Params{
		Limit:  limit,
		Page:   page,
		CatId:  allCategoriesId,
		Search: search,
	}
}

func toIntArray(str string) []int {
	chunks := strings.Split(str, ",")

	var res []int
	for _, c := range chunks {
		i, _ := strconv.Atoi(c) // error handling omitted for concision
		res = append(res, i)
	}

	return res
}
