package password

import (
	"golang.org/x/crypto/bcrypt"
)

type passwordHash struct {
	Cost int
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	NewPassword(password []byte) (string, error)
}

func NewPasswordHash(cost int) PasswordHasher {
	return &passwordHash{Cost: cost}
}

func (p *passwordHash) Hash(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), p.Cost)
	if err != nil {
		return "", err
	}
	return string(passwordHash), nil
}

func (p *passwordHash) NewPassword(password []byte) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword(password, p.Cost)
	if err != nil {
		return "", err
	}
	return string(passwordHash), nil
}
