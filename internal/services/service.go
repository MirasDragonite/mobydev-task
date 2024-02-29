package services

import (
	"miras/internal/models"
	"miras/internal/repository"
)

type Auth interface {
	SignupService(user *models.Register) error
}

type Services struct {
	Auth
}

func NewService(repo *repository.Repository) *Services {
	return &Services{Auth: newAuthService(repo)}
}
