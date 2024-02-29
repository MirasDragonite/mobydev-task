package services

import "miras/internal/repository"

type Services struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Services {
	return &Services{repo: repo}
}
