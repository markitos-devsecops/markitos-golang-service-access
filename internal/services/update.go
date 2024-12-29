package services

import (
	"markitos-golang-service-access/internal/domain"
	"time"
)

type UserUpdateRequestUri struct {
	Id string `uri:"id" binding:"required,uuid"`
}

type UserUpdateRequestBody struct {
	Message string `json:"message" binding:"required"`
}

type UserUpdateRequest struct {
	Id      string `uri:"id" binding:"required,uuid"`
	Message string `json:"message" binding:"required"`
}

func NewUserUpdateRequest(id string, message string) *UserUpdateRequest {
	return &UserUpdateRequest{
		Id:      id,
		Message: message,
	}
}

type UserUpdateService struct {
	Repository domain.UserRepository
}

func NewUserUpdateService(repository domain.UserRepository) UserUpdateService {
	return UserUpdateService{Repository: repository}
}

func (s *UserUpdateService) Execute(request UserUpdateRequest) (*domain.User, error) {
	securedUser, err := domain.NewUser(request.Id, request.Message)
	if err != nil {
		return nil, err
	}

	userToUpdate, errExistingUser := s.Repository.One(&securedUser.Id)
	if errExistingUser != nil {
		return nil, errExistingUser
	}

	userToUpdate.Message = securedUser.Message
	userToUpdate.UpdatedAt = time.Now()
	err = s.Repository.Update(userToUpdate)
	if err != nil {
		return nil, err
	}

	return userToUpdate, nil
}
