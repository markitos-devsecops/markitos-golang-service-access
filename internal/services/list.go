package services

import (
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/domain/dependencies"
)

type UserListService struct {
	Repository dependencies.UserRepository
}

func NewUserListService(repository dependencies.UserRepository) UserListService {
	return UserListService{Repository: repository}
}

func (s *UserListService) Execute() ([]*domain.User, error) {
	response, err := s.Repository.List()
	if err != nil {
		return nil, err
	}

	return response, nil
}
