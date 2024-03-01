package services

import (
	"errors"
	"miras/internal/models"
	"miras/internal/repository"
	"strings"
	"time"
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

func (s *EditService) GetUserData(token string, id int) (models.UserEdit, error) {

	var userData models.UserEdit
	session, err := s.repo.Auth.GetSessionByToken(token)
	if err != nil {
		return userData, err
	}

	user, err := s.repo.Auth.GetUserByID(session.UserId)
	if err != nil {
		return userData, err
	}

	if user.Id != int64(id) {
		return userData, errors.New("Access denied")
	}

	userData.Username = user.Username
	userData.Email = user.Email
	userData.PhoneNum = user.MobilePhone

	userData.BirthDate = user.BirthDate.Format("02 January 2006")

	return userData, nil
}

func (s *EditService) EditUserData(token string, editUser models.UserEdit, id int) error {

	session, err := s.repo.Auth.GetSessionByToken(token)
	if err != nil {
		return err
	}

	user, err := s.repo.Auth.GetUserByID(session.UserId)
	if err != nil {
		return err
	}

	if user.Id != int64(id) {
		return errors.New("Access denied")
	}

	if strings.TrimSpace(editUser.Username) == "" || strings.TrimSpace(editUser.Email) == "" || strings.TrimSpace(editUser.PhoneNum) == "" || strings.TrimSpace(editUser.BirthDate) == "" {
		return errors.New("Empty field")
	}

	user.Username = editUser.Username
	user.Email = editUser.Email
	user.MobilePhone = editUser.PhoneNum
	user.BirthDate, err = time.Parse("02 January 2006", editUser.BirthDate)
	if err != nil {
		return errors.New("Wrong time format (Example:02 January 2006)")
	}
	return s.repo.Edit.EditUserData(&user)
}
