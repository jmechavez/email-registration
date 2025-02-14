package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/jmechavez/email-registration/errors"
	"github.com/jmechavez/email-registration/infrastructure/logger"
	"github.com/jmechavez/email-registration/internal/domain"
)

type UserEmailRepoDb struct {
	emailDb *sqlx.DB
}

// r receiver
func (r UserEmailRepoDb) FindAllUsers() ([]domain.User, *errors.AppError) {
	var err error
	var users []domain.User

	findUserSql := "SELECT id_no, TRIM(COALESCE(first_name, '') || ' ' || COALESCE(last_name, '') || CASE WHEN suffix IS NOT NULL AND suffix != '' THEN ' ' || suffix ELSE '' END) as full_name, email, department, status,srs_no FROM email_accounts"
	err = r.emailDb.Select(&users, findUserSql)
	if err != nil {
		return nil, errors.NewUnExpectedError("Unexpected database error")
	}
	return users, nil
}

func (r UserEmailRepoDb) CreateUserEmail(u domain.User) (*domain.User, *errors.AppError) {
	insertUserSql := "INSERT INTO email_accounts (id_no, first_name, last_name, suffix, email_action,srs_no, department) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING register_no, email"
	err := r.emailDb.QueryRow(
		insertUserSql,
		u.IdNo,
		u.FirstName,
		u.LastName,
		u.Suffix,
		u.EmailAction,
		u.SrsNo,
		u.Department,
	).Scan(&u.RegisterNo, &u.Email)
	if err != nil {
		logger.Error("Error while creating account: " + err.Error())
		return nil, errors.NewUnExpectedError("Unexpected Database Error")
	}
	return &u, nil
}

func NewUserRepoDb(userPostDb *sqlx.DB) UserEmailRepoDb {
	return UserEmailRepoDb{userPostDb}
}
