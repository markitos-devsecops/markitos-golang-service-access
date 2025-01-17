package services_test

import (
	"errors"
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/services"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanGetAUser(t *testing.T) {
	response, _ := userRegisterService.Execute(services.UserRegisterRequest{
		Name:     VALID_NAME,
		Email:    VALID_EMAIL,
		Password: domain.RandomPassword(10),
	})
	require.True(t, domain.IsUUIDv4(response.Id))

	model, err := userMeService.Execute(services.NewUserMeRequest(response.Id))
	require.NoError(t, err)
	require.True(t, domain.IsUUIDv4(model.Id))
	require.True(t, userMockSpyRepository.(*MockSpyUserRepository).OneHaveBeenCalledWithName(model))
}

func TestCantGetOneUserWithEmptyId(t *testing.T) {
	var request services.UserMeRequest = services.UserMeRequest{
		Id: "",
	}

	response, err := userMeService.Execute(request)
	require.Nil(t, response)
	require.Error(t, err)

	var invalidErr *domain.UserInvalidIdError
	require.True(t, errors.As(err, &invalidErr))
	require.Error(t, err)
}
