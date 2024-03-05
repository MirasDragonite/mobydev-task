package repository

import (
	"database/sql"
	"miras/internal/models"
)

type Auth interface {
	CreateUser(user *models.Register) error
	CreateSession(user models.User, token, expiredDate string) error
	GetSessionByToken(token string) (models.Session, error)
	GetUserBy(element, from string) (models.User, error)
	GetUserByID(id int64) (models.User, error)
	UpdateToken(user models.User, token, expaired_data string) error
	DeleteToken(token string) error
	GetSession(userID int64) (models.Session, error)
}

type Edit interface {
	EditUserData(user *models.User) error
}

type Repository struct {
	Auth
	Edit
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Auth: newAuthRepo(db),
		Edit: newEditUserRepository(db),
	}
}
