package service

import (
	"kron-x/internal/dto"
	"kron-x/internal/repository"
	"kron-x/pkg"
	"kron-x/pkg/email"
	"kron-x/pkg/phone"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (s *OrderService) CreateOrder(userId string, orderDto *dto.CreateOrder) error {
	userUUID, err := pkg.UserIdChecker(userId)
	if err != nil {
		return err
	} else {
		err, emailDataUr, emailDataPhyz := s.repo.CreateOrder(userUUID, orderDto)
		if err != nil {
			return err
		} else {
			newEmail := email.NewEmail()
			if emailDataUr == nil {
				if err := newEmail.SendOrderPhyz(emailDataPhyz); err != nil {
					return err
				} else {
					newPhone := phone.NewPhone()
					if err := newPhone.PhoneOrderPhyz(emailDataPhyz); err != nil {
						return err
					} else {
						return nil
					}
				}
			} else {
				if err := newEmail.SendOrderUr(emailDataUr); err != nil {
					return err
				} else {
					return nil
				}
			}
		}
	}
}

func (s *OrderService) GetUserOrder(userId string) ([]*dto.GetOrder, error) {
	userUUID, err := pkg.UserIdChecker(userId)

	if err != nil {
		return nil, err
	} else {
		return s.repo.GetUserOrder(userUUID)
	}
}
