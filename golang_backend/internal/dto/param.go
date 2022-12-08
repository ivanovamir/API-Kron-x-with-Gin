package dto

type Params struct {
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
	CatId  []int  `gorm:"type:int" json:"cat_id"`
	Search string `json:"search"`
}
