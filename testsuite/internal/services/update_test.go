package services_test

import (
	"errors"
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/services"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanUpdateAUser(t *testing.T) {
	createdUser, err := userCreateService.Execute(services.UserCreateRequest{
		Name: VALID_NAME,
	})
	require.NoError(t, err)

	model, err := userUpdateService.Execute(services.UserUpdateRequest{
		Id:   createdUser.Id,
		Name: createdUser.Name + " updated",
	})

	require.NoError(t, err)
	require.Equal(t, createdUser.Id, model.Id)
	require.NotEqual(t, createdUser.Name, model.Name)
	require.Equal(t, createdUser.Name+" updated", model.Name)

	require.True(t, userMockSpyRepository.(*MockSpyUserRepository).UpdateHaveBeenCalledWithName(model))
}

func TestCantUpdatOneUserWithEmptyName(t *testing.T) {
	var request services.UserUpdateRequest = services.UserUpdateRequest{
		Id:   VALID_UUIDV4,
		Name: "",
	}

	response, err := userUpdateService.Execute(request)
	require.Nil(t, response)
	require.Error(t, err)

	var invalidErr *domain.UserInvalidNameError
	require.True(t, errors.As(err, &invalidErr))
	require.Error(t, err)
}

func TestCantUpdatOneUserWithEmptyId(t *testing.T) {
	var request services.UserUpdateRequest = services.UserUpdateRequest{
		Id:   "",
		Name: VALID_NAME,
	}

	response, err := userUpdateService.Execute(request)
	require.Nil(t, response)
	require.Error(t, err)

	var invalidErr *domain.UserInvalidIdError
	require.True(t, errors.As(err, &invalidErr))
	require.Error(t, err)
}

func TestCantUpdatOneUserWithInvalidId(t *testing.T) {
	var request services.UserUpdateRequest = services.UserUpdateRequest{
		Id:   "invalid-id",
		Name: VALID_NAME,
	}

	response, err := userUpdateService.Execute(request)
	require.Nil(t, response)
	require.Error(t, err)

	var invalidErr *domain.UserInvalidIdFormatError
	require.True(t, errors.As(err, &invalidErr))
	require.Error(t, err)
}
