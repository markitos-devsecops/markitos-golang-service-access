package services

import (
	"markitos-golang-service-access/internal/domain"
)

type UserCreateRequest struct {
	Username string `json:"username" binding:"required"`
}

func NewBolilerCreateRequest(username string) UserCreateRequest {
	return UserCreateRequest{Username: username}
}

type UserCreateService struct {
	Repository domain.UserRepository
}

func NewUserCreateService(repository domain.UserRepository) UserCreateService {
	return UserCreateService{Repository: repository}
}

func (s *UserCreateService) Execute(request UserCreateRequest) (*domain.User, error) {
	user, err := domain.NewUser(domain.UUIDv4(), request.Username)
	if err != nil {
		return nil, err
	}

	err = s.Repository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
