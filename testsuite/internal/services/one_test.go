package services_test

import (
	"errors"
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/services"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanGetAUser(t *testing.T) {
	response, _ := userCreateService.Execute(services.UserCreateRequest{
		Name: VALID_NAME,
	})
	require.True(t, domain.IsUUIDv4(response.Id))

	model, err := userOneService.Execute(services.NewUserOneRequest(response.Id))
	require.NoError(t, err)
	require.True(t, domain.IsUUIDv4(model.Id))
	require.True(t, userMockSpyRepository.(*MockSpyUserRepository).OneHaveBeenCalledWithName(model))
}

func TestCantGetOneUserWithEmptyId(t *testing.T) {
	var request services.UserOneRequest = services.UserOneRequest{
		Id: "",
	}

	response, err := userOneService.Execute(request)
	require.Nil(t, response)
	require.Error(t, err)

	var invalidErr *domain.UserInvalidIdError
	require.True(t, errors.As(err, &invalidErr))
	require.Error(t, err)
}
