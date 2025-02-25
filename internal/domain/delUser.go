package domain

import (
	"github.com/jmechavez/email-registration/errors"
	"github.com/jmechavez/email-registration/internal/dto"
)

type DelUser struct {
	IdNo        string `json:"id_no"        db:"id_no"`
	FirstName   string `json:"first_name"   db:"first_name"`
	LastName    string `json:"last_name"    db:"last_name"`
	Suffix      string `json:"suffix"       db:"suffix"`
	FullName    string `json:"full_name"    db:"full_name"`
	Department  string `json:"department"   db:"department"`
	Email       string `json:"email"        db:"email"`
	Status      string `json:"status"       db:"status"`
	EmailAction string `json:"email_action" db:"email_action"`
	SrsNo       string `json:"srs_no"       db:"srs_no"`
}

func (du DelUser) ToDelUserResponseDto() dto.DeleteUserEmailResponse {
	return dto.DeleteUserEmailResponse{
		IdNo:       du.IdNo,
		Department: du.Department,
		FullName:   du.FullName,
	}
}

type DelUserRepo interface {
	DeleteUserEmail(DelUser) (*DelUser, *errors.AppError)
}
