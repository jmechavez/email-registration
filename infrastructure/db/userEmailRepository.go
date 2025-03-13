package db

import (
	"database/sql"

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
	findAllUserSql := "SELECT id_no, TRIM(COALESCE(first_name, '') || ' ' || COALESCE(last_name, '') || CASE WHEN suffix IS NOT NULL AND suffix != '' THEN ' ' || suffix ELSE '' END) as full_name, email, department, status, srs_no FROM email_accounts"
	if err := r.emailDb.Select(&users, findAllUserSql); err != nil {
		return nil, errors.NewUnExpectedError("Unexpected database error")
	}
	return users, nil
}

func (r UserEmailRepoDb) IdNo(idNo string) (*domain.User, *errors.AppError) {
	// First check active email accounts
	findActiveEmailSql := `
		SELECT id_no, first_name, last_name, suffix, email, email_action, date_created, srs_no, department
		FROM email_accounts
		WHERE id_no = $1
	`
	var user domain.User
	err := r.emailDb.QueryRow(findActiveEmailSql, idNo).Scan(
		&user.IdNo,
		&user.FirstName,
		&user.LastName,
		&user.Suffix,
		&user.Email,
		&user.EmailAction,
		&user.DateCreated,
		&user.SrsNo,
		&user.Department,
	)

	// If found active account, return it
	if err == nil {
		// Set DateDeleted to empty string or nil depending on your struct
		return &user, nil
	}

	// If not found or other error, check deleted accounts
	if err != sql.ErrNoRows {
		logger.Error("Error while querying active email accounts: " + err.Error())
		return nil, errors.NewUnExpectedError("Unexpected database error")
	}

	// Check deleted email accounts
	findDeletedEmailSql := `
		SELECT id_no, first_name, last_name, suffix, email, email_action, date_created, srs_no, department, deleted_at
		FROM deleted_email_accounts
		WHERE id_no = $1
	`

	var deletedUser domain.User
	// If your DateDeleted is a string
	err = r.emailDb.QueryRow(findDeletedEmailSql, idNo).Scan(
		&deletedUser.IdNo,
		&deletedUser.FirstName,
		&deletedUser.LastName,
		&deletedUser.Suffix,
		&deletedUser.Email,
		&deletedUser.EmailAction,
		&deletedUser.DateCreated,
		&deletedUser.SrsNo,
		&deletedUser.Department,
		&deletedUser.DateDeleted,
	)

	if err == nil {
		return &deletedUser, nil
	}

	if err == sql.ErrNoRows {
		return nil, errors.NewNotFoundError("Client not found")
	}

	logger.Error("Error while querying deleted email accounts: " + err.Error())
	return nil, errors.NewUnExpectedError("Unexpected database error")
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

// NewUserRepoDb initializes the UserEmailRepoDb repository.
func NewUserRepoDb(userPostDb *sqlx.DB) UserEmailRepoDb {
	return UserEmailRepoDb{userPostDb}
}
