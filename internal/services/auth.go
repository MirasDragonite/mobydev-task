package services

import (
	"errors"
	"miras/internal/models"
	"miras/internal/repository"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
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

func (s *AuthService) SigninService(data *models.Login) (http.Cookie, error) {

	if data.Passowrd != data.Repassword {
		return http.Cookie{}, errors.New("bad request")
	}

	user, err := s.repo.Auth.GetUserBy(data.Email, "email")
	if err != nil {
		return http.Cookie{}, err
	}
	if !CheckPasswordHash(data.Passowrd, user.HashedPassword) {
		return http.Cookie{}, errors.New("Passwords din't match")
	}

	token, err := createToken(user.HashedPassword)
	if err != nil {
		return http.Cookie{}, err
	}

	cookie := http.Cookie{
		Name:     "Token",
		Expires:  time.Now().Add(time.Hour * 24),
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	}
	session := models.Session{
		UserId:      user.Id,
		Token:       token,
		ExpiredDate: time.Now().Add(time.Minute * 30).Format("2006-01-02 15:04:05"),
	}
	err = s.repo.Auth.CreateSession(&session)
	if err != nil {
		return http.Cookie{}, err
	}
	return cookie, nil

}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

var secretKey = []byte("secret-key")

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Minute * 30).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
