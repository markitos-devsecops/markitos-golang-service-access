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
		Message: VALID_MESSAGE,
	})
	require.NoError(t, err)

	model, err := userUpdateService.Execute(services.UserUpdateRequest{
		Id:      createdUser.Id,
		Message: createdUser.Message + " updated",
	})

	require.NoError(t, err)
	require.Equal(t, createdUser.Id, model.Id)
	require.NotEqual(t, createdUser.Message, model.Message)
	require.Equal(t, createdUser.Message+" updated", model.Message)

	require.True(t, userMockSpyRepository.(*MockSpyUserRepository).UpdateHaveBeenCalledWithMessage(model))
}

func TestCantUpdatOneUserWithEmptyMessage(t *testing.T) {
	var request services.UserUpdateRequest = services.UserUpdateRequest{
		Id:      VALID_UUIDV4,
		Message: "",
	}

	response, err := userUpdateService.Execute(request)
	require.Nil(t, response)
	require.Error(t, err)

	var invalidErr *domain.UserInvalidMessageError
	require.True(t, errors.As(err, &invalidErr))
	require.Error(t, err)
}

func TestCantUpdatOneUserWithEmptyId(t *testing.T) {
	var request services.UserUpdateRequest = services.UserUpdateRequest{
		Id:      "",
		Message: VALID_MESSAGE,
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
		Id:      "invalid-id",
		Message: VALID_MESSAGE,
	}

	response, err := userUpdateService.Execute(request)
	require.Nil(t, response)
	require.Error(t, err)

	var invalidErr *domain.UserInvalidIdFormatError
	require.True(t, errors.As(err, &invalidErr))
	require.Error(t, err)
}
