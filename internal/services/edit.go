package services

import (
	"miras/internal/models"
	"miras/internal/repository"
)

type EditService struct {
	repo *repository.Repository
}

func newEditService(repo *repository.Repository) *EditService {
	return &EditService{repo: repo}
}

func (s *EditService) GetUserByToken(token string) (models.User, error) {

	session, err := s.repo.Auth.GetSessionByToken(token)
	if err != nil {
		return models.User{}, err
	}

	user, err := s.repo.Auth.GetUserByID(session.UserId)
	if err != nil {
		return models.User{}, err
	}

	return user, nil

}

func (s *EditService) EditUserData(user models.User) error {

	return s.repo.Edit.EditUserData(&user)
}
