package dto

import "time"

type OptionOfDeliveryAndPayment struct {
	City            string `json:"city"`
	Street          string `json:"street"`
	House           string `json:"house"`
	Frame           string `json:"frame"`
	Entrance        string `json:"entrance"`
	ApartmentOffice string `json:"appartment_office"`
	OrderID         uint   `json:"order_id"`
}

type CreateOrder struct {
	Note            string                      `json:"note"`
	DeliveryPayment *OptionOfDeliveryAndPayment `json:"delivery_payment"`
}

type GetOrder struct {
	Id        string    `json:"id"`
	Note      string    `json:"note"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
