package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"kron-x/internal/dto"
	"kron-x/internal/repository"
	"kron-x/pkg/email"
	"kron-x/pkg/phone"
	"math/rand"
	"os"
	"time"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

func CodeCreator() int { //4 digit unique code
	rand.Seed(time.Now().UnixNano())
	code := 1000 + rand.Intn(9999-1000)
	return code
}

type CustomClaims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userId int) string {
	claims := CustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(4320 * time.Hour)), //TODO mb change expires time
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNED_KEY")))
	return tokenString
}

func (s *UserService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv("SIGNED_KEY")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *CustomClaims")
	}

	return claims.UserId, nil
}

func (s *UserService) RegisterUser(userDTO *dto.RegisterUser) (string, error) {
	newUserId, err := s.repo.RegisterUser(userDTO)
	if err != nil {
		return "", err
	} else {
		token := GenerateToken(newUserId)
		if userDTO.Email != "" {
			newEmail := email.NewEmail()
			if err := newEmail.SendEmailPassword(userDTO.Email, userDTO.Password); err != nil {
				return "", err
			} else {
				return token, nil
			}
		}
		return token, nil
	}
}

func (s *UserService) AuthenticateUser(userDTO *dto.UserAuth) (string, error) {
	userId, err := s.repo.AuthenticateUser(userDTO)
	if err != nil {
		return "", err
	} else {
		token := GenerateToken(userId)
		return token, nil
	}
}

func (s *UserService) ChangeUserPassword(userInfoDto *dto.UserInformation) error {
	newPassword, err := s.repo.ChangeUserPassword(userInfoDto)
	if err != nil {
		return err
	} else {
		newEmail := email.NewEmail()
		if err := newEmail.SenEmailChangePassword(userInfoDto.Email, fmt.Sprint(newPassword)); err != nil {
			return err
		} else {
			return nil
		}
	}
}

func (s *UserService) CodeGenerator(codeGenerator *dto.CodeGenerator) error {
	code := CodeCreator()
	if err := s.repo.CodeGenerator(codeGenerator, code); err != nil {
		return err
	} else {
		if codeGenerator.Email == "" {
			newPhone := phone.NewPhone()
			if err := newPhone.SendCode(fmt.Sprint(code), codeGenerator.Phone); err != nil {
				return err
			} else {
				return nil
			}
		}
		newEmail := email.NewEmail()
		if err := newEmail.SendEmailCode(codeGenerator.Email, fmt.Sprint(code)); err != nil {
			return err
		} else {
			return nil
		}
	}
}
