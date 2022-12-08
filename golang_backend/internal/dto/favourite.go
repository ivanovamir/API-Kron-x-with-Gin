package dto

type UpdateFavourite struct {
	ProductId string `json:"product_id" binding:"required"`
}

type DeleteFavourite struct {
	ProductId string `json:"product_id" binding:"required"`
}
