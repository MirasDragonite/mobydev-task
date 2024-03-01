package repository

import (
	"database/sql"
	"miras/internal/models"
)

type EditUserRepository struct {
	db *sql.DB
}

func newEditUserRepository(db *sql.DB) *EditUserRepository {
	return &EditUserRepository{db: db}
}

func (r *EditUserRepository) EditUserData(user *models.User) error {

	query := `UPDATE users SET username=$1,email=$2,mobile_phone=$3,birth_date=$4 WHERE id=$5`

	_, err := r.db.Exec(query, user.Username, user.Email, user.MobilePhone, user.BirthDate.Format("02 January 2006"), user.Id)

	if err != nil {
		return err
	}

	return nil
}
