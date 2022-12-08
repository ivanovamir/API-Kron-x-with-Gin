package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "kron-x/docs"
	"kron-x/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(
		CORSMiddleware(),
	)

	auth := router.Group("/auth")
	{
		auth.POST("/sign_up", h.RegisterUser)
		auth.POST("/sign_in", h.AuthenticateUser)
		auth.POST("/new_password", h.ChangeUserPassword)
		auth.POST("/code", h.CodeGenerator)
	}

	api := router.Group("/api", h.UserIdentity)
	{
		cart := api.Group("/cart")
		{
			cart.GET("/", h.GetUserCart)
			cart.POST("/add", h.UpdateUserCart)
			cart.POST("/del", h.DeleteProduct)
		}
		favourite := api.Group("/favourite")
		{
			favourite.GET("/", h.GetUserFavourites)
			favourite.POST("/add", h.UpdateUserFavourites)
			favourite.POST("/del", h.DeleteFavouriteProduct)

		}
		order := api.Group("/order")
		{
			order.POST("/", h.CreateOrder)
			order.GET("/", h.GetUserOrder)
		}
	}

	categories := router.Group("/category")
	{
		categories.GET("/", h.GetAllCategories)
	}
	router.GET("/category_detail", h.GetCategoriesById)
	products := router.Group("/products")
	{
		products.GET("/", h.GetAllProductsByParams)

	}
	router.GET("/product_detail", h.GetProductDetail)
	cms := router.Group("/cms")
	{
		cms.GET("/sliders", h.GetAllSliders)
		cms.GET("/about", h.GetAbout)
		cms.GET("/req", h.GetRequisite)
		cms.GET("/headoffice", h.GetAllHeadOffice)
		cms.GET("/features", h.GetAllFeature)
		cms.GET("/service", h.GetAllService)
		cms.GET("/service_detail", h.GetServiceById)
		cms.POST("/feedback_call", h.FeedbackCall)
		cms.POST("/feedback_selection", h.FeedbackSelection)
		cms.POST("/support", h.Support)
	}
	search := router.Group("/search")
	{
		search.GET("", h.GetSearch)
	}
	catalog := router.Group("/catalog")
	{
		catalog.GET("", h.GetAllCatalogs)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router

}
