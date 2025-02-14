package service

import (
	"github.com/jmechavez/email-registration/errors"
	"github.com/jmechavez/email-registration/internal/domain"
	"github.com/jmechavez/email-registration/internal/dto"
)

type UserService interface {
	GetAllUserEmails() ([]dto.UserEmailResponse, *errors.AppError)
	NewUser(dto.NewUserEmailRequest) (*dto.NewUserEmailResponse, *errors.AppError)
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

func (r DefaultUserService) NewUser(
	req dto.NewUserEmailRequest,
) (*dto.NewUserEmailResponse, *errors.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	user := domain.User{
		RegisterNo:  "",
		IdNo:        req.IdNo,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Suffix:      req.Suffix,
		Department:  req.Department,
		Email:       "",
		Status:      "active",
		EmailAction: req.EmailAction,
		SrsNo:       req.SrsNo,
	}
	newUser, err := r.repo.CreateUserEmail(user)
	if err != nil {
		return nil, err
	}
	response := newUser.ToNewUserResponseDto()

	return &response, nil
}

func NewUserService(repository domain.UserRepo) DefaultUserService {
	return DefaultUserService{repository}
}
