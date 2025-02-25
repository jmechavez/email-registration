package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/jmechavez/email-registration/errors"
	"github.com/jmechavez/email-registration/infrastructure/logger"
	"github.com/jmechavez/email-registration/internal/domain"
)

type DelUserEmailRepoDb struct {
	emailDb *sqlx.DB
}

// DeleteUserEmail removes a user email record from the database by ID and assigns a new srs_no.
func (r DelUserEmailRepoDb) DeleteUserEmail(
	du domain.DelUser,
) (*domain.DelUser, *errors.AppError) {
	deleteUserSql := `
		WITH moved AS (
			INSERT INTO deleted_email_accounts (id_no, srs_no, first_name, last_name, suffix, email, email_action, date_created, department)
			SELECT id_no, $2, first_name, last_name, COALESCE(suffix, ''), email, 'delete', date_created, department
			FROM email_accounts
			WHERE id_no = $1
			RETURNING id_no, first_name, last_name, NULLIF(suffix, '') AS suffix
		)
		DELETE FROM email_accounts
		WHERE id_no IN (SELECT id_no FROM moved)
		RETURNING 
			id_no, 
			department,
			first_name, 
			last_name, 
			COALESCE(NULLIF(suffix, ''), 'None') AS suffix, 
			TRIM(
				COALESCE(first_name, '') || ' ' || COALESCE(last_name, '') || 
				CASE WHEN suffix IS NOT NULL AND suffix != '' THEN ' ' || suffix ELSE '' END
			) AS full_name;
	`
	var deletedUser domain.DelUser

	// Fix: Pass individual struct fields instead of the whole struct
	row := r.emailDb.QueryRow(deleteUserSql, du.IdNo, du.SrsNo)

	// Fix: Scan values into the deletedUser struct
	err := row.Scan(
		&deletedUser.IdNo,
		&deletedUser.Department,
		&deletedUser.FirstName,
		&deletedUser.LastName,
		&deletedUser.Suffix,
		&deletedUser.FullName,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned, meaning the id_no does not exist
			logger.Error("User not found with id_no: " + du.IdNo)
			return nil, errors.NewNotFoundError("User not found")
		} else {
			// Other database errors
			logger.Error("Error while deleting user: " + err.Error())
			return nil, errors.NewUnExpectedError("Unexpected Database Error")
		}
	}

	return &deletedUser, nil
}

func DelUserRepoDb(userDeleteDb *sqlx.DB) DelUserEmailRepoDb {
	return DelUserEmailRepoDb{userDeleteDb}
}
