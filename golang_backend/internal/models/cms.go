package models

import (
	"time"
)

type Slider struct {
	Id        uint `gorm:"primaryKey"`
	MainText  string
	UpperText string
	DownText  string
	Link      string
	Image     string
}

type About struct {
	Id             uint `gorm:"primaryKey"`
	MainImage      string
	SecondaryImage string
	Image1         string
	Image2         string
	Image3         string
	Image4         string
	Image5         string
	Image6         string
	Image7         string
}

type Metric struct {
	Id           uint `gorm:"primaryKey"`
	GoogleMetric string
	YandexMetric string
}
type Requisite struct {
	Id          uint `gorm:"primaryKey"`
	CompanyName string
	Address     string
	Inn         string
	Kpp         string
	Ogrn        string
	Rs          string
	Bank        string
	Ks          string
	Bik         string
}

type HeadOffice struct {
	Id       uint `gorm:"primaryKey"`
	Title    string
	Phone    string
	Mail     string
	Schedule string
	Address  string
}

type Feature struct {
	Id     uint `gorm:"primaryKey"`
	Header string
	Body   string
	Icon   string
}

type Service struct {
	Id        uint `gorm:"primary_key"`
	Title     string
	Body      string
	Image     string
	CreatedAt time.Time
}

type Email struct {
	Id         uint `gorm:"primaryKey"`
	LogoImage  string
	CartImage  string
	CheckImage string
	Address    string
}

type FeedbackCall struct {
	Id    uint `gorm:"primary_key"`
	Name  string
	Phone string
}

type FeedbackSelection struct {
	Id      uint `gorm:"primaryKey"`
	Name    string
	Phone   string
	Email   string
	Comment string
}

type Support struct {
	Id   uint `gorm:"primaryKey"`
	Name string
	Body string
}
