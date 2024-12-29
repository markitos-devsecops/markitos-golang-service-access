package services_test

import (
	"errors"
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/services"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanCreateAUser(t *testing.T) {
	var request services.UserCreateRequest = services.UserCreateRequest{
		Username: VALID_USERNAME,
	}
	response, err := userCreateService.Execute(request)

	require.NoError(t, err)
	require.True(t, domain.IsUUIDv4(response.Id))
	require.Equal(t, VALID_USERNAME, response.Username)
	require.True(t, userMockSpyRepository.(*MockSpyUserRepository).CreateHaveBeenCalledWithUsername(response))
}

func TestCantCreateAUserWithEmptyMessage(t *testing.T) {
	var request services.UserCreateRequest = services.UserCreateRequest{
		Username: "",
	}

	response, err := userCreateService.Execute(request)
	require.Nil(t, response)
	require.Error(t, err)

	var invalidErr *domain.UserInvalidUsernameError
	require.True(t, errors.As(err, &invalidErr))
	require.Error(t, err)
}
