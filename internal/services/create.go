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
	Repository dependencies.UserRepository
}

func NewUserCreateService(repository dependencies.UserRepository) UserCreateService {
	return UserCreateService{Repository: repository}
}

func (s *UserCreateService) Execute(request UserCreateRequest) (*domain.User, error) {
	var hashedPassword string = request.Password
	user, err := domain.NewUser(domain.UUIDv4(), request.Name, request.Email, hashedPassword)
	if err != nil {
		return nil, err
	}

	err = s.Repository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
