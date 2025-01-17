package services

import (
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/domain/dependencies"
	"time"
)

type UserUpdateMeRequestUri struct {
	Id string `uri:"id" binding:"required,uuid"`
}

type UserUpdateMeRequestBody struct {
	Name string `json:"name" binding:"required"`
}

type UserUpdateMeRequest struct {
	Id   string `uri:"id" binding:"required,uuid"`
	Name string `json:"name" binding:"required"`
}

func NewUserUpdateMeRequest(id string, name string) *UserUpdateMeRequest {
	return &UserUpdateMeRequest{
		Id:   id,
		Name: name,
	}
}

type UserUpdateMeService struct {
	Repository dependencies.UserRepository
}

func NewUserUpdateMeService(repository dependencies.UserRepository) UserUpdateMeService {
	return UserUpdateMeService{Repository: repository}
}

func (s *UserUpdateMeService) Execute(request UserUpdateMeRequest) (*domain.User, error) {
	securedId, err := domain.NewUserId(request.Id)
	if err != nil {
		return nil, err
	}
	securedName, err := domain.NewUserName(request.Name)
	if err != nil {
		return nil, err
	}

	var id string = securedId.Value()
	userToUpdate, errExistingUser := s.Repository.One(&id)
	if errExistingUser != nil {
		return nil, errExistingUser
	}

	userToUpdate.Name = securedName.Value()
	userToUpdate.UpdatedAt = time.Now()
	err = s.Repository.Update(userToUpdate)
	if err != nil {
		return nil, err
	}

	return userToUpdate, nil
}
