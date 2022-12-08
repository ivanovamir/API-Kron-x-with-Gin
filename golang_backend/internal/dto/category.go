package dto

type Category struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Image      string `json:"image"`
	CountItems string `json:"countitems"`
}

type CategoryParams struct {
	ID string
}
