package repository

import (
	"gorm.io/gorm"
	"kron-x/internal/dto"
)

type Repository struct {
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
	RegisterUser(userDTO *dto.RegisterUser) (int, error)
	AuthenticateUser(userDTO *dto.UserAuth) (int, error)
	ChangeUserPassword(userInfoDto *dto.UserInformation) (int, error)
	CodeGenerator(codeGenerator *dto.CodeGenerator, code int) error
}

type Cart interface {
	GetUserCart(userUUID int) ([]*dto.CartProduct, error)
	UpdateUserCart(userUUID int, product *dto.UpdateCart) error
	DeleteCartProduct(userUUID int, productId int) error
}

type Favourite interface {
	GetUserFavourites(userUUID int) ([]*dto.FavouritesProducts, error)
	UpdateUserFavourites(userUUID int, product *dto.UpdateFavourite) error
	DeleteFavouriteProduct(userUUID int, productId int) error
}

type Order interface {
	CreateOrder(userId int, orderDto *dto.CreateOrder) (error, *dto.SendOrderEmailUr, *dto.SendOrderEmailPhyz)
	GetUserOrder(userId int) ([]*dto.GetOrder, error)
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

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Category:  NewCategoryRepository(db),
		Product:   NewProductRepository(db),
		CMS:       NewCMSRepository(db),
		User:      NewUserRepository(db),
		Cart:      NewCartRepository(db),
		Favourite: NewFavouriteRepository(db),
		Order:     NewOrderRepository(db),
		Search:    NewSearchRepository(db),
		Catalog:   NewCatalogRepository(db),
	}
}
