package services

import (
	"errors"
	"miras/internal/models"
	"miras/internal/repository"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.Repository
}

func newAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignupService(user *models.Register) error {
	if strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.Passowrd) == "" {
		return errors.New("Empty data")
	}
	hashedPswd, err := HashPassword(user.Passowrd)
	if err != nil {
		return err
	}
	user.Passowrd = hashedPswd
	return s.repo.CreateUser(user)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
