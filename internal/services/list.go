package services

import (
	"markitos-golang-service-access/internal/domain"
)

type UserListService struct {
	Repository domain.UserRepository
}

func NewUserListService(repository domain.UserRepository) UserListService {
	return UserListService{Repository: repository}
}

func (s *UserListService) Execute() ([]*domain.User, error) {
	response, err := s.Repository.List()
	if err != nil {
		return nil, err
	}

	return response, nil
}
