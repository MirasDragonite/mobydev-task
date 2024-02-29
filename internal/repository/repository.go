package repository

import (
	"database/sql"
	"miras/internal/models"
)

type Auth interface {
	CreateUser(user *models.Register) error
	CreateSession(session *models.Session) error
	GetSessionByToken(token string) (models.Session, error)
	GetUserBy(element, from string) (models.User, error)
	GetUserByID(id int64) (models.User, error)
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
