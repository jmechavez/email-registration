package service

import (
	"github.com/jmechavez/email-registration/errors"
	"github.com/jmechavez/email-registration/internal/domain"
	"github.com/jmechavez/email-registration/internal/dto"
)

type UserService interface {
	GetAllUserEmails() ([]dto.UserEmailResponse, *errors.AppError)
}

type DefaultUserService struct {
	repo domain.UserRepo
}

func (r DefaultUserService) GetAllUserEmails() ([]dto.UserEmailResponse, *errors.AppError) {
	u, err := r.repo.FindAllUsers()
	if err != nil {
		return nil, err
	}

	var uResponse []dto.UserEmailResponse
	for _, user := range u {
		uResponse = append(uResponse, user.ToDto())
	}

	return uResponse, nil
}

func NewUserService(repository domain.UserRepo) DefaultUserService {
	return DefaultUserService{repository}
}
