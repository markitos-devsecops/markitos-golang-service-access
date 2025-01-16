package services

import (
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/domain/dependencies"
	"time"
)

type UserUpdateRequestUri struct {
	Id string `uri:"id" binding:"required,uuid"`
}

type UserUpdateRequestBody struct {
	Name string `json:"name" binding:"required"`
}

type UserUpdateRequest struct {
	Id   string `uri:"id" binding:"required,uuid"`
	Name string `json:"name" binding:"required"`
}

func NewUserUpdateRequest(id string, name string) *UserUpdateRequest {
	return &UserUpdateRequest{
		Id:   id,
		Name: name,
	}
}

type UserUpdateService struct {
	Repository dependencies.UserRepository
}

func NewUserUpdateService(repository dependencies.UserRepository) UserUpdateService {
	return UserUpdateService{Repository: repository}
}

func (s *UserUpdateService) Execute(request UserUpdateRequest) (*domain.User, error) {
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
