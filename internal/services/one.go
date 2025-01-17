package services

import (
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/domain/dependencies"
)

type UserMeRequest struct {
	Id string `uri:"id" binding:"required" minLength:"36" maxLength:"36"`
}

func NewUserMeRequest(id string) UserMeRequest {
	return UserMeRequest{Id: id}
}

type UserMeService struct {
	Repository dependencies.UserRepository
}

func NewUserMeService(repository dependencies.UserRepository) UserMeService {
	return UserMeService{Repository: repository}
}

func (s *UserMeService) Execute(request UserMeRequest) (*domain.User, error) {
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
