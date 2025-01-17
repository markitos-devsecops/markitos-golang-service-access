package services

import (
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/domain/dependencies"
)

type UserRegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserRegisterService struct {
	Repository     dependencies.UserRepository
	PasswordHasher dependencies.Hasher
}

func NewUserRegisterService(repository dependencies.UserRepository, hasher dependencies.Hasher) UserRegisterService {
	return UserRegisterService{Repository: repository, PasswordHasher: hasher}
}

func (s *UserRegisterService) Execute(request UserRegisterRequest) (*domain.User, error) {
	user, err := domain.NewUser(domain.UUIDv4(), request.Name, request.Email, request.Password)
	if err != nil {
		return nil, err
	}
	hashedPassword, err := s.PasswordHasher.Create(request.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	err = s.Repository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
