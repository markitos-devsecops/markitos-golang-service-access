package services

import (
	"markitos-golang-service-access/internal/domain"
)

type UserCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

func NewBolilerCreateRequest(name string) UserCreateRequest {
	return UserCreateRequest{Name: name}
}

type UserCreateService struct {
	Repository domain.UserRepository
}

func NewUserCreateService(repository domain.UserRepository) UserCreateService {
	return UserCreateService{Repository: repository}
}

func (s *UserCreateService) Execute(request UserCreateRequest) (*domain.User, error) {
	user, err := domain.NewUser(domain.UUIDv4(), request.Name)
	if err != nil {
		return nil, err
	}

	err = s.Repository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
