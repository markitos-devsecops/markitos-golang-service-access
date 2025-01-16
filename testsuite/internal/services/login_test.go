package services_test

import (
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/services"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanLoginExistingUser(t *testing.T) {
	var request services.UserCreateRequest = services.UserCreateRequest{
		Name:     domain.RandomPersonName(),
		Email:    VALID_EMAIL,
		Password: domain.RandomPassword(10),
	}
	createdUser, err := userCreateService.Execute(request)
	require.NoError(t, err)

	model, err := userLoginService.Execute(services.UserLoginRequest{
		Email:    request.Email,
		Password: request.Password,
	})
	require.NoError(t, err)
	require.NotNil(t, createdUser)
	require.NotNil(t, model)
	// require.True(t, userMockSpyRepository.(*MockSpyUserRepository).LoginHaveBeenCalled())
	// require.True(t, userMockSpyHasher.(*MockSpyUserHasher).CreateHaveBeenCalled())
	// require.True(t, userMockSpyTokener.(*MockSpyTokener).CreateHaveBeenCalled())
}
