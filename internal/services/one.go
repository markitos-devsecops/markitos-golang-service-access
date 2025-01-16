package services

import (
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/domain/dependencies"
)

type UserOneRequest struct {
	Id string `uri:"id" binding:"required" minLength:"36" maxLength:"36"`
}

func NewUserOneRequest(id string) UserOneRequest {
	return UserOneRequest{Id: id}
}

type UserOneService struct {
	Repository dependencies.UserRepository
}

func NewUserOneService(repository dependencies.UserRepository) UserOneService {
	return UserOneService{Repository: repository}
}

func (s *UserOneService) Execute(request UserOneRequest) (*domain.User, error) {
	requestedId := &request.Id
	userId, err := domain.NewUserId(*requestedId)
	if err != nil {
		return nil, err
	}

	secureId := userId.Value()
	response, err := s.Repository.One(&secureId)
	if err != nil {
		return nil, err
	}

	return response, nil
}
