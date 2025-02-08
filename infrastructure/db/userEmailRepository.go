package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/jmechavez/email-registration/errors"
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

func NewUserRepoDb(userPostDb *sqlx.DB) UserEmailRepoDb {
	return UserEmailRepoDb{userPostDb}
}
