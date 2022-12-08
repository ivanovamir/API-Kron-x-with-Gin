package dto

import "time"

type Slider struct {
	Id        string `json:"id"`
	MainText  string `json:"main_text"`
	UpperText string `json:"upper_text"`
	DownText  string `json:"down_text"`
	Link      string `json:"url_link"`
	Image     string `json:"image"`
}

type About struct {
	MainImage      string `json:"main_image"`
	SecondaryImage string `json:"secondary_image"`
	Image1         string `json:"main_office_image_1"`
	Image2         string `json:"main_office_image_2"`
	Image3         string `json:"main_office_image_3"`
	Image4         string `json:"main_office_image_4"`
	Image5         string `json:"main_office_image_5"`
	Image6         string `json:"stock_image_6"`
	Image7         string `json:"stock_image_7"`
}

type Metric struct {
	Id           string `json:"id" `
	GoogleMetric string `json:"google_metric" `
	YandexMetric string `json:"yandex_metric"`
}
type Requisites struct {
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Inn         string `json:"inn"`
	Kpp         string `json:"kpp"`
	Ogrn        string `json:"ogrn"`
	Rs          string `json:"rs"`
	Bank        string `json:"bank"`
	Ks          string `json:"ks"`
	Bik         string `json:"bik"`
}

type HeadOffice struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Phone    string `json:"phone"`
	Mail     string `json:"mail"`
	Schedule string `json:"schedule"`
	Address  string `json:"address"`
}

type Feature struct {
	Id     string `json:"id"`
	Header string `json:"header"`
	Body   string `json:"body"`
	Icon   string `json:"icon"`
}

type Email struct {
	Id         string `json:"id"`
	LogoImage  string `json:"logo_image"`
	CartImage  string `json:"cart_image"`
	CheckImage string `json:"check_image"`
	Address    string `json:"address"`
}

type Service struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
}

type FeedbackCall struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}

type FeedbackSelection struct {
	Name    string `json:"name" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Email   string `json:"mail" binding:"required"`
	Comment string `json:"comment" binding:"required"`
}

type Support struct {
	Name string `json:"name" binding:"required"`
	Body string `json:"body" binding:"required"`
}
