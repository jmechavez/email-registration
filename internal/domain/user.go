package domain

import (
	"github.com/jmechavez/email-registration/errors"
	"github.com/jmechavez/email-registration/internal/dto"
)

type User struct {
	RegisterNo  string `json:"register_no"  db:"register_no"`
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

func (u User) ToDto() dto.UserEmailResponse {
	return dto.UserEmailResponse{
		RegisterNo:  u.RegisterNo,
		IdNo:        u.IdNo,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Suffix:      u.Suffix,
		FullName:    u.FullName,
		Department:  u.Department,
		Email:       u.Email,
		Status:      u.Status,
		EmailAction: u.EmailAction,
		SrsNo:       u.SrsNo,
	}
}

func (u User) ToNewUserResponseDto() dto.NewUserEmailResponse {
	return dto.NewUserEmailResponse{
		RegisterNo: u.RegisterNo,
		Email:      u.Email,
	}
}

type UserRepo interface {
	FindAllUsers() ([]User, *errors.AppError)
	CreateUserEmail(User) (*User, *errors.AppError)
}
