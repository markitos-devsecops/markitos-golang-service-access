package domain_test

import (
	"errors"
	"markitos-golang-service-access/internal/domain"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUserWithEmptyId(t *testing.T) {
	user, err := domain.NewUser("", domain.RandomPersonName(), domain.RandomEmail(), domain.RandomPassword(10))

	var invalidIdErr *domain.UserInvalidIdError
	require.True(t, errors.As(err, &invalidIdErr))
	require.Equal(t, domain.NewUserInvalidIdError("").Error(), err.Error())
	require.True(t, strings.HasPrefix(err.Error(), domain.USER_INVALID_ID))
	require.Error(t, err)
	require.Empty(t, user)
}

func TestCreateUserWithEmptyName(t *testing.T) {
	user, err := domain.NewUser(VALID_UUIDV4, "", domain.RandomEmail(), domain.RandomPassword(10))

	var invalidErr *domain.UserInvalidNameError
	require.True(t, errors.As(err, &invalidErr))
	require.Equal(t, domain.NewUserInvalidNameError("").Error(), err.Error())
	require.True(t, strings.HasPrefix(err.Error(), domain.USER_INVALID_NAME))
	require.Error(t, err)
	require.Empty(t, user)
}
