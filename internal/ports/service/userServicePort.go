package service

import (
	"github.com/jmechavez/email-registration/errors"
	"github.com/jmechavez/email-registration/internal/domain"
	"github.com/jmechavez/email-registration/internal/dto"
)

type UserService interface {
	GetAllUserEmails() ([]dto.UserEmailResponse, *errors.AppError)
	NewUser(dto.NewUserEmailRequest) (*dto.NewUserEmailResponse, *errors.AppError)
	GetUserByIdNo(idNo string) (*dto.ByIdNoUserEmailResponse, *errors.AppError)
}

type DelUserService interface {
	DelUser(dto.DeleteUserEmailRequest) (*dto.DeleteUserEmailResponse, *errors.AppError)
}

type DefaultDelUserService struct {
	repo domain.DelUserRepo
}

type DefaultUserService struct {
	repo domain.UserRepo
}

func (r DefaultUserService) GetUserByIdNo(idNo string) (*dto.ByIdNoUserEmailResponse, *errors.AppError) {
	if idNo == "" {
		return nil, errors.NewValidationError("Id number cannot be empty")
	}

	user, err := r.repo.IdNo(idNo)
	if err != nil {
		return nil, err
	}

	response := user.ByIdNoUserEmailResponse()
	return &response, nil
}

func (r DefaultDelUserService) DelUser(req dto.DeleteUserEmailRequest) (*dto.DeleteUserEmailResponse, *errors.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	delUser := domain.DelUser{
		IdNo:        req.IdNo,
		EmailAction: "deleted",
		SrsNo:       req.SrsNo,
	}
	deletedUser, err := r.repo.DeleteUserEmail(delUser)
	if err != nil {
		return nil, err
	}
	response := deletedUser.ToDelUserResponseDto()
	return &response, nil
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

func (r DefaultUserService) NewUser(req dto.NewUserEmailRequest) (*dto.NewUserEmailResponse, *errors.AppError) {
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

func NewDelUserService(repository domain.DelUserRepo) DefaultDelUserService {
	return DefaultDelUserService{repository}
}

func NewUserService(repository domain.UserRepo) DefaultUserService {
	return DefaultUserService{repository}
}
