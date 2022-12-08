package phone

import (
	"kron-x/internal/dto"
	"os"
)

type Phone interface {
	SendCode(code, phone string) error
	PhoneOrderPhyz(smsOrderPhyz *dto.SendOrderEmailPhyz) error
}

func NewPhone() Phone {
	return &phone{
		ApiKey: os.Getenv("API_KEY"),
	}
}
