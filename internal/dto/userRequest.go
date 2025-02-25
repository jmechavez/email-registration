package dto

import "github.com/jmechavez/email-registration/errors"

type NewUserEmailRequest struct {
	IdNo        string `json:"id_no"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Suffix      string `json:"suffix"`
	Department  string `json:"department"`
	SrsNo       string `json:"srs_no"`
	EmailAction string `json:"email_action"`
}

type DeleteUserEmailRequest struct {
	IdNo        string `json:"id_no"`
	SrsNo       string `json:"srs_no"`
	EmailAction string `json:"email_action"`
}

func (r NewUserEmailRequest) Validate() *errors.AppError {
	if r.EmailAction != "save" {
		return errors.NewValidationError("User action failed")
	}
	return nil
}

func (r DeleteUserEmailRequest) Validate() *errors.AppError {
	return nil
}
