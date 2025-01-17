package services

import (
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/domain/dependencies"
)

type UserCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func NewBolilerCreateRequest(name string) UserCreateRequest {
	return UserCreateRequest{Name: name}
}

type UserCreateService struct {
	Repository     dependencies.UserRepository
	PasswordHasher dependencies.Hasher
}

func NewUserCreateService(repository dependencies.UserRepository, hasher dependencies.Hasher) UserCreateService {
	return UserCreateService{Repository: repository, PasswordHasher: hasher}
}

func (s *UserCreateService) Execute(request UserCreateRequest) (*domain.User, error) {
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
