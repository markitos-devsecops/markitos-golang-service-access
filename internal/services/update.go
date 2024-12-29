package services

import (
	"markitos-golang-service-access/internal/domain"
	"time"
)

type UserUpdateRequestUri struct {
	Id string `uri:"id" binding:"required,uuid"`
}

type UserUpdateRequestBody struct {
}

type UserUpdateRequest struct {
	Id string `uri:"id" binding:"required,uuid"`
}

func NewUserUpdateRequest(id string) *UserUpdateRequest {
	return &UserUpdateRequest{
		Id: id,
	}
}

type UserUpdateService struct {
	Repository domain.UserRepository
}

func NewUserUpdateService(repository domain.UserRepository) UserUpdateService {
	return UserUpdateService{Repository: repository}
}

func (s *UserUpdateService) Execute(request UserUpdateRequest) (*domain.User, error) {
	secureId, err := domain.NewUserId(request.Id)
	if err != nil {
		return nil, err
	}

	value := secureId.Value()
	existingUser, err := s.Repository.One(&value)
	if err != nil {
		return nil, err
	}

	existingUser.UpdatedAt = time.Now()
	err = s.Repository.Update(existingUser)
	if err != nil {
		return nil, err
	}

	return existingUser, nil
}
