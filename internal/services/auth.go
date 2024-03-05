package services

import (
	"database/sql"
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

	newToken, err := createToken(user.HashedPassword)
	if err != nil {
		return http.Cookie{}, err
	}
	cookie := http.Cookie{
		Name:     "Token",
		Path:     "/",
		HttpOnly: true,
	}

	newSessonTime := time.Now().Add(time.Minute * 24)
	session, err := s.repo.Auth.GetSession(user.Id)

	if err != nil {
		if err == sql.ErrNoRows {

			cookie.Value = newToken
			cookie.Expires = newSessonTime
			err = s.repo.Auth.CreateSession(user, cookie.Value, cookie.Expires.Format("2006-01-02 15:04:05"))
			if err != nil {
				return http.Cookie{}, err
			}
			return cookie, nil

		} else {
			return http.Cookie{}, err
		}
	} else {
		sessionTime, err := time.Parse("2006-01-02 15:04:05", session.ExpiredDate)
		if err != nil {
			return http.Cookie{}, err
		}
		if sessionTime.After(time.Now()) || sessionTime.Equal(time.Now()) {

			err = s.repo.UpdateToken(user, newToken, time.Now().Format("2006-01-02 15:04:05"))
			if err != nil {
				return http.Cookie{}, err
			}
			cookie.Value = newToken
			cookie.Expires = time.Now()
		} else {
			cookie.Value = session.Token
			cookie.Expires = sessionTime
		}
	}

	return cookie, nil
}

func (s *AuthService) DeleteToken(cookie *http.Cookie) error {
	err := s.repo.DeleteToken(cookie.Value)

	cookie.Name = "Token"
	cookie.Value = ""
	cookie.Path = "/"
	cookie.MaxAge = -1
	cookie.HttpOnly = false
	if err != nil {
		return err
	}
	return nil
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
