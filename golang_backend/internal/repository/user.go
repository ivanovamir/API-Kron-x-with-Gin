package repository

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"kron-x/internal/dto"
	models2 "kron-x/internal/models"
	"kron-x/pkg/password"
	"math/rand"
	"time"
)

const (
	cost        = 7
	userCtx     = "userId"
	passwordNew = 9999999999999999
)

var (
	ErrDB             = fmt.Errorf("error uccured in db")
	ErrUserRegistered = fmt.Errorf("user already registered")
	ErrWrongCode      = fmt.Errorf("wrong code")
	ErrUnregistered   = fmt.Errorf("unregistered")
	ErrUnauthorized   = fmt.Errorf("unauthorized")
	ErrWrongPassword  = fmt.Errorf("invalid password")
	ErrUserNotFound   = fmt.Errorf("user not found")
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) RegisterUser(userDTO *dto.RegisterUser) (int, error) {
	var user, newUser *models2.User
	var code *models2.Code

	switch r.db.Where("email = ?", userDTO.Email).Last(&code).Error {
	case nil:
		switch r.db.Where("email = ?", userDTO.Email).Or("inn = ?", userDTO.Inn).First(&user).Error {
		case nil:
			return 0, ErrUserRegistered
		default:
			if userDTO.Code != code.Code {
				return 0, ErrWrongCode
			} else {
				cost := password.NewPasswordHash(cost)
				passwordHash, err := cost.Hash(userDTO.Password)

				if err != nil {
					panic(err)
				}

				newUser = &models2.User{
					CreatedAt:    time.Now(),
					Inn:          userDTO.Inn,
					Email:        userDTO.Email,
					PasswordHash: passwordHash,
				}
				r.db.Where("email LIKE ?", userDTO.Email).Unscoped().Delete(&models2.Code{})
			}
		}
	default:
		return 0, ErrUnregistered
	}

	r.db.Create(&newUser)

	createdCart := &models2.Cart{UserID: newUser.Id}
	r.db.Create(&createdCart)

	createdFavouriteList := &models2.Favourites{UserID: newUser.Id}
	r.db.Create(&createdFavouriteList)

	return int(newUser.Id), nil
}

func (r *UserRepository) AuthenticateUser(userDTO *dto.UserAuth) (int, error) {
	var user, newUser *models2.User
	var code *models2.Code

	//TODO We dont neet to check for empty email or phone,
	switch userDTO.Email {
	case "":
		switch r.db.Where("phone = ?", userDTO.Phone).Last(&code).Error {
		case nil:
			if userDTO.Code != code.Code {
				return 0, ErrWrongCode
			} else {
				r.db.Where("phone LIKE ?", userDTO.Phone).Unscoped().Delete(&models2.Code{})
				switch r.db.Where("phone = ?", userDTO.Phone).First(&user) {
				case nil:
					return int(user.Id), nil
				default:
					newUser = &models2.User{
						CreatedAt: time.Now(),
						Phone:     userDTO.Phone,
					}
					r.db.Create(&newUser)

					createdCart := &models2.Cart{UserID: newUser.Id}
					r.db.Create(&createdCart)

					createdFavouriteList := &models2.Favourites{UserID: newUser.Id}
					r.db.Create(&createdFavouriteList)
					return int(newUser.Id), nil
				}
			}
		default:
			return 0, ErrUnauthorized
		}
	default:
		switch r.db.Where("email = ?", userDTO.Email).First(&user).Error {
		case nil:
			switch bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(userDTO.Password)) {
			case nil:
				return int(user.Id), nil
			default:
				return 0, ErrWrongPassword
			}
		default:
			return 0, ErrUnauthorized
		}
	}
}

func (r *UserRepository) ChangeUserPassword(userInfoDto *dto.UserInformation) (int, error) {
	var user *models2.User
	switch r.db.Where("email = ?", userInfoDto.Email).First(&user).Error {
	case nil:
		rand.Seed(time.Now().UnixNano())
		generatedPassword := rand.Intn(passwordNew)
		passwordByte := []byte(fmt.Sprint(generatedPassword))
		newCost := password.NewPasswordHash(cost)

		passwordHash, err := newCost.NewPassword(passwordByte)
		if err != nil {
			panic(err)
		}
		user.PasswordHash = passwordHash
		r.db.Save(&user)
		return generatedPassword, nil
	default:
		return 0, ErrUserNotFound
	}
}

func (r *UserRepository) CodeGenerator(codeGenerator *dto.CodeGenerator, code int) error {
	var newCode *models2.Code
	var user *models2.User

	if codeGenerator.Email == "" {
		r.db.Where("phone = ?", codeGenerator.Phone).First(&user)
		newCode = &models2.Code{
			Phone: codeGenerator.Phone,
			Code:  fmt.Sprint(code),
		}
		if err := r.db.Create(&newCode).Error; err != nil {
			return ErrDB
		} else {
			return nil
		}
	}

	switch r.db.Where("email = ?", codeGenerator.Email).First(&user).Error {
	case nil:
		return ErrUserRegistered
	default:
		newCode = &models2.Code{
			Email: codeGenerator.Email,
			Code:  fmt.Sprint(code),
		}
		if err := r.db.Create(&newCode).Error; err != nil {
			return ErrDB
		} else {
			return nil
		}
	}
}
