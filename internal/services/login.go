package services

import (
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/domain/dependencies"
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
	Repository dependencies.UserRepository
}

func NewUserLoginService(
	repository dependencies.UserRepository,
	hasher dependencies.Hasher,
	tokener dependencies.Tokener) UserLoginService {
	return UserLoginService{Repository: repository}
}

func (s *UserLoginService) Execute(request UserLoginRequest) (*domain.User, error) {
	securedEmail, err := domain.NewUserEmail(request.Email)
	if err != nil {
		return nil, err
	}
	securedPassword, err := domain.NewUserPassword(request.Password)
	if err != nil {
		return nil, err
	}

	user, _ := domain.NewUser(domain.UUIDv4(), domain.RandomPersonName(), securedEmail.Value(), securedPassword.Value())

	return user, nil
}
