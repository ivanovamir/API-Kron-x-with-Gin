package repository

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"kron-x/internal/dto"
	models2 "kron-x/internal/models"
	"time"
)

var (
	ErrCartEmpty     = fmt.Errorf("cart is empty")
	ErrOrderNotFound = fmt.Errorf("order not found")
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) CreateOrder(userId int, orderDto *dto.CreateOrder) (error, *dto.SendOrderEmailUr, *dto.SendOrderEmailPhyz) {

	var newOrder *models2.Order
	var cart *models2.Cart
	var newCart *models2.Cart
	var cartProducts []*models2.CartProducts
	var emailData *models2.Email
	var orderEmailUr *dto.SendOrderEmailUr
	var orderEmailPhyz *dto.SendOrderEmailPhyz
	var user *models2.User
	var orderProducts []*dto.OrderProduct
	var totalPrice float64

	r.db.Last(&emailData)

	r.db.Where("id = ?", userId).First(&user)

	switch r.db.Where("in_order = ?", false).Where("user_id = ?", userId).First(&cart).Error {
	case nil:
		r.db.Preload(clause.Associations).
			Joins("inner join cartm2ms ug on ug.cart_products_id = cart_products.id ").
			Joins("inner join carts g on g.id= ug.cart_id ").
			Where("g.in_order = false AND g.user_id = ?", uint(userId)).Find(&cartProducts)

		switch orderDto.DeliveryPayment {
		case nil:
			newOrder = &models2.Order{
				Note:      orderDto.Note,
				UserID:    uint(userId),
				CartID:    cart.Id,
				CreatedAt: time.Now(),
			}
		default:
			newOrder = &models2.Order{
				Note: orderDto.Note,
				OptionOfDeliveryAndPayment: &models2.OptionOfDeliveryAndPayment{
					City:            orderDto.DeliveryPayment.City,
					Street:          orderDto.DeliveryPayment.Street,
					House:           orderDto.DeliveryPayment.House,
					Frame:           orderDto.DeliveryPayment.Frame,
					Entrance:        orderDto.DeliveryPayment.Entrance,
					ApartmentOffice: orderDto.DeliveryPayment.ApartmentOffice,
				},
				UserID:    uint(userId),
				CartID:    cart.Id,
				CreatedAt: time.Now(),
			}
		}

		cart.InOrder = true
		r.db.Save(&cart)

		r.db.Create(&newOrder)

		newCart = &models2.Cart{
			UserID: uint(userId),
		}
		r.db.Create(&newCart)

		if user.Phone == "" {
			for x := range cartProducts {
				totalPrice += cartProducts[x].Product.Price * float64(cartProducts[x].Count)

				orderProd := &dto.OrderProduct{
					Title: cartProducts[x].Product.Title,
					Count: fmt.Sprint(cartProducts[x].Count),
					Price: fmt.Sprint(cartProducts[x].Product.Price),
				}

				orderProducts = append(orderProducts, orderProd)
			}

			orderEmailUr = &dto.SendOrderEmailUr{
				Address:    emailData.Address,
				OrderId:    fmt.Sprint(newOrder.Id),
				Email:      user.Email,
				Inn:        user.Inn,
				Comment:    newOrder.Note,
				LogoImage:  emailData.LogoImage,
				CartImage:  emailData.CartImage,
				CheckImage: emailData.CheckImage,
				Product:    orderProducts,
				TotalPrice: fmt.Sprint(totalPrice),
			}
			return nil, orderEmailUr, nil
		} else {
			for x := range cartProducts {
				totalPrice += cartProducts[x].Product.Price * float64(cartProducts[x].Count)

				orderProd := &dto.OrderProduct{
					Title: cartProducts[x].Product.Title,
					Count: fmt.Sprint(cartProducts[x].Count),
					Price: fmt.Sprint(cartProducts[x].Product.Price),
				}

				orderProducts = append(orderProducts, orderProd)
			}

			orderEmailPhyz = &dto.SendOrderEmailPhyz{
				Address:    emailData.Address,
				OrderId:    fmt.Sprint(newOrder.Id),
				Phone:      user.Phone,
				Comment:    newOrder.Note,
				LogoImage:  emailData.LogoImage,
				CartImage:  emailData.CartImage,
				CheckImage: emailData.CheckImage,
				Product:    orderProducts,
				TotalPrice: fmt.Sprint(totalPrice),
			}
			return nil, nil, orderEmailPhyz

		}
	default:
		return ErrCartEmpty, nil, nil
	}
}

func (r *OrderRepository) GetUserOrder(userId int) ([]*dto.GetOrder, error) {
	var order []*models2.Order
	var allOrdersDto []*dto.GetOrder

	switch r.db.Where("user_id = ?", userId).Find(&order).Error {
	case nil:
		if len(order) == 0 {
			return []*dto.GetOrder{}, nil
		} else {
			for x := range order {
				var orderStatus *models2.OrderStatus
				r.db.Where("id = ?", order[x].OrderStatusID).First(&orderStatus)

				orderDto := &dto.GetOrder{
					Id:        fmt.Sprint(order[x].Id),
					Note:      order[x].Note,
					Status:    orderStatus.Name,
					CreatedAt: order[x].CreatedAt,
				}
				allOrdersDto = append(allOrdersDto, orderDto)
			}
			return allOrdersDto, nil
		}

	default:
		return nil, ErrOrderNotFound
	}
}
