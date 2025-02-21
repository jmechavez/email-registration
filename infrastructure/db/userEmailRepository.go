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

// FindAllUsers retrieves all user email records from the database.
func (r UserEmailRepoDb) FindAllUsers() ([]domain.User, *errors.AppError) {
	var users []domain.User
	findUserSql := "SELECT id_no, TRIM(COALESCE(first_name, '') || ' ' || COALESCE(last_name, '') || CASE WHEN suffix IS NOT NULL AND suffix != '' THEN ' ' || suffix ELSE '' END) as full_name, email, department, status, srs_no FROM email_accounts"
	if err := r.emailDb.Select(&users, findUserSql); err != nil {
		return nil, errors.NewUnExpectedError("Unexpected database error")
	}
	return users, nil
}

// CreateUserEmail inserts a new user email record into the database.
func (r UserEmailRepoDb) CreateUserEmail(u domain.User) (*domain.User, *errors.AppError) {
	insertUserSql := "INSERT INTO email_accounts (id_no, first_name, last_name, suffix, email_action, srs_no, department) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING register_no, email"
	if err := r.emailDb.QueryRow(
		insertUserSql,
		u.IdNo,
		u.FirstName,
		u.LastName,
		u.Suffix,
		u.EmailAction,
		u.SrsNo,
		u.Department,
	).Scan(&u.RegisterNo, &u.Email); err != nil {
		logger.Error("Error while creating account: " + err.Error())
		return nil, errors.NewUnExpectedError("Unexpected Database Error")
	}
	return &u, nil
}

// DeleteUserEmail removes a user email record from the database by ID and assigns a new srs_no.
func (r UserEmailRepoDb) DeleteUserEmail(
	du domain.DelUser,
) (*domain.DelUser, *errors.AppError) {
	deleteUserSql := `WITH moved AS (
    INSERT INTO deleted_email_accounts (id_no, srs_no, first_name, last_name, suffix, email, email_action, date_created, department)
    SELECT id_no, $2, first_name, last_name, COALESCE(suffix, ''), email, 'delete', date_created, department
    FROM email_accounts WHERE id_no = $1
    RETURNING id_no, first_name, last_name, NULLIF(suffix, '') AS suffix
    )
    DELETE FROM email_accounts WHERE id_no IN (SELECT id_no FROM moved)
    RETURNING id_no, 
          department,
          first_name, 
          last_name, 
          COALESCE(NULLIF(suffix, ''), 'None') AS suffix, 
          TRIM(COALESCE(first_name, '') || ' ' || COALESCE(last_name, '') || 
               CASE WHEN suffix IS NOT NULL AND suffix != '' THEN ' ' || suffix ELSE '' END) AS full_name;`

	var deletedUser domain.DelUser

	// Fix: Pass individual struct fields instead of the whole struct
	row := r.emailDb.QueryRow(deleteUserSql, du.IdNo, du.SrsNo)

	// Fix: Scan values into the deletedUser struct
	err := row.Scan(
		&deletedUser.IdNo,
		&deletedUser.FirstName,
		&deletedUser.LastName,
		&deletedUser.Suffix,
	)
	if err != nil {
		logger.Error("Error while deleting user: " + err.Error())
		return nil, errors.NewUnExpectedError("Unexpected Database Error")
	}

	return &deletedUser, nil
}

// NewUserRepoDb initializes the UserEmailRepoDb repository.
func NewUserRepoDb(userPostDb *sqlx.DB) UserEmailRepoDb {
	return UserEmailRepoDb{userPostDb}
}
