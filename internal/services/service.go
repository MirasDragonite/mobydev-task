package services

import (
	"miras/internal/models"
	"miras/internal/repository"
	"net/http"
)

type Auth interface {
	SignupService(user *models.Register) error
	SigninService(data *models.Login) (http.Cookie, error)
	DeleteToken(cookie *http.Cookie) error
}

type Edit interface {
	GetUserByToken(token string) (models.User, error)
	GetUserData(token string, id int) (models.UserEdit, error)
	EditUserData(token string, editUser models.UserEdit, id int) error
}

type Services struct {
	Auth
	Edit
}

func NewService(repo *repository.Repository) *Services {
	return &Services{
		Auth: newAuthService(repo),
		Edit: newEditService(repo),
	}
}
