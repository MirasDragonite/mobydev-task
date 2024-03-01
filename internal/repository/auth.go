package repository

import (
	"database/sql"
	"fmt"
	"miras/internal/models"
	"strings"
	"time"
)

type AuthRepository struct {
	db *sql.DB
}

func newAuthRepo(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user *models.Register) error {
	query := `INSERT INTO users(username,email,hash_password,mobile_phone,birth_date) VALUES($1,$2,$3,$4,$5)`

	_, err := r.db.Exec(query, "", user.Email, user.Passowrd, "", "")
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthRepository) GetUserBy(element, from string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT * FROM users WHERE %s=$1 ", from)
	row := r.db.QueryRow(query, element)
	birthdDate := ""

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.HashedPassword, &user.MobilePhone, &birthdDate)
	if err != nil {

		return models.User{}, err
	}
	if strings.TrimSpace(birthdDate) != "" {
		user.BirthDate, err = time.Parse("02 January 2006", birthdDate)
		if err != nil {

			return models.User{}, err
		}
	}
	return user, nil
}

func (r *AuthRepository) CreateSession(session *models.Session) error {
	query := `INSERT INTO sessions(user_id,token,expired_date) VALUES($1,$2,$3)`

	_, err := r.db.Exec(query, session.UserId, session.Token, session.ExpiredDate)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthRepository) GetSessionByToken(token string) (models.Session, error) {
	var session models.Session

	query := `SELECT * FROM sessions WHERE token=$1 `
	row := r.db.QueryRow(query, token)
	err := row.Scan(&session.Id, &session.UserId, &session.Token, &session.ExpiredDate)
	if err != nil {
		return models.Session{}, err
	}

	return session, nil
}

func (r *AuthRepository) GetUserByID(id int64) (models.User, error) {
	var user models.User
	birthdDate := ""
	query := `SELECT * FROM users WHERE id=$1 `
	row := r.db.QueryRow(query, id)
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.HashedPassword, &user.MobilePhone, &birthdDate)
	if err != nil {
		return models.User{}, err
	}
	if strings.TrimSpace(birthdDate) != "" {
		user.BirthDate, err = time.Parse("02 January 2006", birthdDate)
		if err != nil {

			return models.User{}, err
		}
	}

	return user, nil
}
