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
}

type Repository struct{ Auth }

func NewRepository(db *sql.DB) *Repository {
	return &Repository{Auth: newAuthRepo(db)}
}
