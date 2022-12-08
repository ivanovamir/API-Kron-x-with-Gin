package dto

type UpdateCart struct {
	ProductId string `json:"product_id" binding:"required"`
	Count     string `json:"count" binding:"required"`
}

type DeleteCart struct {
	ProductId string `json:"product_id" binding:"required"`
}
