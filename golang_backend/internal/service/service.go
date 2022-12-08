package service

import (
	"kron-x/internal/dto"
	"kron-x/internal/repository"
)

type Service struct {
	Category
	Product
	CMS
	User
	Cart
	Favourite
	Order
	Search
	Catalog
}

type Category interface {
	GetAllCategories() ([]*dto.Category, error)
	GetCategoriesById(id *dto.CategoryParams) ([]*dto.Category, error)
}

type Product interface {
	GetAllProductsByParams(params *dto.Params) ([]*dto.Product, error)
	GetProductDetail(id string) (*dto.ProductInformation, error)
	ParseUrlParams() []int
}

type User interface {
	ParseToken(accessToken string) (int, error)
	RegisterUser(userDTO *dto.RegisterUser) (string, error)
	AuthenticateUser(userDTO *dto.UserAuth) (string, error)
	ChangeUserPassword(userInfoDto *dto.UserInformation) error
	CodeGenerator(codeGenerator *dto.CodeGenerator) error
}

type Cart interface {
	GetUserCart(userId string) ([]*dto.CartProduct, error)
	UpdateUserCart(userId string, product *dto.UpdateCart) error
	DeleteProduct(userId string, productId int) error
}

type Favourite interface {
	GetUserFavourites(userId string) ([]*dto.FavouritesProducts, error)
	UpdateUserFavourites(userId string, product *dto.UpdateFavourite) error
	DeleteFavouriteProduct(userId string, productId int) error
}

type Order interface {
	CreateOrder(userId string, orderDto *dto.CreateOrder) error
	GetUserOrder(userId string) ([]*dto.GetOrder, error)
}

type CMS interface {
	GetAllSliders() ([]*dto.Slider, error)
	GetAbout() ([]*dto.About, error)
	GetRequisite() ([]*dto.Requisites, error)
	GetAllHeadOffice() ([]*dto.HeadOffice, error)
	GetAllFeature() ([]*dto.Feature, error)
	GetAllService() ([]*dto.Service, error)
	GetServiceById(id string) ([]*dto.Service, error)
	FeedbackCall(feedbackCallData *dto.FeedbackCall) error
	FeedbackSelection(feedbackSelectionData *dto.FeedbackSelection) error
	Support(supportData *dto.Support) error
}

type Search interface {
	GetSearch(searchParams *dto.Params) ([]*dto.Product, error)
}

type Catalog interface {
	GetAllCatalogs(params *dto.Params) ([]*dto.Catalog, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Category:  NewCategoryService(repo),
		Product:   NewProductService(repo),
		CMS:       NewCMSService(repo),
		User:      NewUserService(repo),
		Cart:      NewCartService(repo),
		Favourite: NewFavouriteService(repo),
		Order:     NewOrderService(repo),
		Search:    NewSearchService(repo),
		Catalog:   NewCatalogService(repo),
	}
}
