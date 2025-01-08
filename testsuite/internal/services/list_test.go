package services_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanListAUsers(t *testing.T) {
	models, err := userListService.Execute()

	require.NoError(t, err)
	require.NotNil(t, models)
	require.True(t, userMockSpyRepository.(*MockSpyUserRepository).ListHaveBeenCalled())
}
