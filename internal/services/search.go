package services

import (
	"markitos-golang-service-access/internal/domain"
)

type UserSearchService struct {
	Repository domain.UserRepository
}

func NewUserSearchService(repository domain.UserRepository) UserSearchService {
	return UserSearchService{Repository: repository}
}

type UserSearchRequest struct {
	SearchTerm string `json:"searchTerm"`
	PageNumber int    `json:"pageNumber" bindings:"min=1"`
	PageSize   int    `json:"pageSize" bindings:"min=10,max=100"`
}

func (s *UserSearchService) Execute(request UserSearchRequest) ([]*domain.User, error) {
	response, err := s.Repository.SearchAndPaginate(request.SearchTerm, request.PageNumber, request.PageSize)
	if err != nil {
		return nil, err
	}

	return response, nil
}
