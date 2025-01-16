package services

import (
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/domain/dependencies"
	"time"
)

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func NewUserLoginRequest(email string, password string) *UserLoginRequest {
	return &UserLoginRequest{
		Email:    email,
		Password: password,
	}
}

type UserLoginService struct {
	Repository     dependencies.UserRepository
	PasswordHasher dependencies.Hasher
	UserTokener    dependencies.Tokener
}

func NewUserLoginService(
	repository dependencies.UserRepository,
	hasher dependencies.Hasher,
	tokener dependencies.Tokener) UserLoginService {

	return UserLoginService{
		Repository:     repository,
		PasswordHasher: hasher,
		UserTokener:    tokener,
	}
}

func (s *UserLoginService) Execute(request UserLoginRequest) (string, error) {
	securedEmail, err := domain.NewUserEmail(request.Email)
	if err != nil {
		return "", err
	}
	securedPassword, err := domain.NewUserPassword(request.Password)
	if err != nil {
		return "", err
	}

	user, _ := s.Repository.OneFromEmailAndPassword(securedEmail.Value(), securedPassword.Value())
	if user == nil {
		return "", domain.NewUnauthorizedError()
	}

	if !s.PasswordHasher.Validate(user.Password, securedPassword.Value()) {
		return "", domain.NewUnauthorizedError()
	}

	token, err := s.UserTokener.Create(user.Id, time.Hour*24)
	if err != nil {
		return "", err
	}

	return token, nil
}
