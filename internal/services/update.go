package services

import (
	"markitos-golang-service-access/internal/domain"
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
	Repository domain.UserRepository
}

func NewUserUpdateService(repository domain.UserRepository) UserUpdateService {
	return UserUpdateService{Repository: repository}
}

func (s *UserUpdateService) Execute(request UserUpdateRequest) (*domain.User, error) {
	securedUser, err := domain.NewUser(request.Id, request.Name)
	if err != nil {
		return nil, err
	}

	userToUpdate, errExistingUser := s.Repository.One(&securedUser.Id)
	if errExistingUser != nil {
		return nil, errExistingUser
	}

	userToUpdate.Name = securedUser.Name
	userToUpdate.UpdatedAt = time.Now()
	err = s.Repository.Update(userToUpdate)
	if err != nil {
		return nil, err
	}

	return userToUpdate, nil
}
