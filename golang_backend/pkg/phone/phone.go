package phone

import (
	"fmt"
	"kron-x/internal/dto"
	"net/http"
	"os"
)

const (
	brandName = "Kron-x"
	phoneUrl  = "http://api.sms-prosto.ru/?method=push_msg&key=%s&text=%s&phone=%s&sender_name=%s"
)

type phone struct {
	ApiKey string
}

func (p *phone) SendCode(code, phone string) error {

	message := fmt.Sprintf("Ваш код - %s", code)

	url := fmt.Sprintf(phoneUrl, os.Getenv("API_KEY"), message, phone, brandName)

	resp, err := http.Get(url)

	if err != nil {
		return err
	} else {
		defer resp.Body.Close()
		return nil
	}
}

func (p *phone) PhoneOrderPhyz(smsOrderPhyz *dto.SendOrderEmailPhyz) error {

	var messageData string

	for x := range smsOrderPhyz.Product {
		messageData += fmt.Sprint(smsOrderPhyz.Product[x].Title, " ", smsOrderPhyz.Product[x].Count, " ", "шт.", " ", smsOrderPhyz.Product[x].Price, " ", "руб", " | ")
	}
	message := fmt.Sprintf("Ваш заказ с сайта Kron-x Номер заказа - %s Товары | Кол-во | Цена: %s Итоговая цена: %s Адресс: %s", fmt.Sprint(smsOrderPhyz.OrderId), messageData, fmt.Sprint(smsOrderPhyz.TotalPrice), smsOrderPhyz.Address)

	fmt.Println(message)

	url := fmt.Sprintf(phoneUrl, os.Getenv("API_KEY"), message, smsOrderPhyz.Phone, brandName)
	resp, err := http.Get(url)
	if err != nil {
		return err
	} else {
		defer resp.Body.Close()
		return nil
	}
}
