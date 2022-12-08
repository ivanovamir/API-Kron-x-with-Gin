package dto

type Product struct {
	Id            string `json:"id"`
	Category      string `json:"category"`
	CategoryName  string `json:"category_name"`
	Title         string `json:"title"`
	VendorCode    string `json:"vendor_code"`
	Code1c        string `json:"code_1c"`
	Description   string `json:"description"`
	Price         string `json:"price"`
	ImageOriginal string `json:"image_original"`
	Image128      string `json:"image_128"`
	Image432      string `json:"image_432"`
}

type ProductInformation struct {
	Id            string `json:"id"`
	Category      string `json:"category"`
	CategoryName  string `json:"category_name"`
	Title         string `json:"title"`
	VendorCode    string `json:"vendor_code"`
	Code1c        string `json:"code_1c"`
	Quantity      string `json:"quantity"`
	Description   string `json:"description"`
	Price         string `json:"price"`
	ImageOriginal string `json:"image_original"`
	Image128      string `json:"image_128"`
	Image432      string `json:"image_432"`
}

type CartProduct struct {
	ID           string `json:"id"`
	Category     string `json:"category"`
	CategoryName string `json:"category_name"`
	Title        string `json:"title"`
	VendorCode   string `json:"vendor_code"`
	Price        string `json:"price"`
	Image128     string `json:"image_128"`
	Count        string `json:"count"`
}

type FavouritesProducts struct {
	ID           string `json:"id"`
	Category     string `json:"category"`
	CategoryName string `json:"category_name"`
	Title        string `json:"title"`
	VendorCode   string `json:"vendor_code"`
	Price        string `json:"price"`
	Image128     string `json:"image_128"`
}

type OrderProduct struct {
	Title string
	Count string
	Price string
}
