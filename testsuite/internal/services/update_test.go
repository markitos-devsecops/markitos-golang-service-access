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
		Username: domain.RandomEmail(),
	})

	require.NoError(t, err)

	model, err := userUpdateService.Execute(services.UserUpdateRequest{
		Id: createdUser.Id,
	})

	require.NoError(t, err)
	require.Equal(t, createdUser.Id, model.Id)
	require.NotEqual(t, createdUser.Username, model.Username)

	require.True(t, userMockSpyRepository.(*MockSpyUserRepository).UpdateHaveBeenCalledWithUsername(model))
}

func TestCantUpdatOneUserWithEmptyId(t *testing.T) {
	var request services.UserUpdateRequest = services.UserUpdateRequest{
		Id: "",
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
		Id: "invalid-id",
	}

	response, err := userUpdateService.Execute(request)
	require.Nil(t, response)
	require.Error(t, err)

	var invalidErr *domain.UserInvalidIdFormatError
	require.True(t, errors.As(err, &invalidErr))
	require.Error(t, err)
}
