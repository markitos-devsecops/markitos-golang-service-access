package services_test

import (
	"errors"
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/services"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanCreateAUser(t *testing.T) {
	var request services.UserRegisterRequest = services.UserRegisterRequest{
		Name:     VALID_NAME,
		Email:    VALID_EMAIL,
		Password: domain.RandomPassword(10),
	}
	response, err := userRegisterService.Execute(request)

	require.NoError(t, err)
	require.True(t, domain.IsUUIDv4(response.Id))
	require.Equal(t, VALID_NAME, response.Name)
	require.True(t, userMockSpyRepository.(*MockSpyUserRepository).CreateHaveBeenCalledWithName(response))
}

func TestCantCreateAUserWithEmptyName(t *testing.T) {
	var request services.UserRegisterRequest = services.UserRegisterRequest{
		Name:     "",
		Email:    VALID_EMAIL,
		Password: domain.RandomPassword(10),
	}
	response, err := userRegisterService.Execute(request)
	require.Nil(t, response)
	require.Error(t, err)

	var invalidErr *domain.UserInvalidNameError
	require.True(t, errors.As(err, &invalidErr))
	require.Error(t, err)
}
