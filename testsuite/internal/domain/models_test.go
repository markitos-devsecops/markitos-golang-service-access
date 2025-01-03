package domain_test

import (
	"errors"
	"markitos-golang-service-access/internal/domain"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUserWithEmptyId(t *testing.T) {
	user, err := domain.NewUser("", domain.RandomEmail())

	var invalidIdErr *domain.UserInvalidIdError
	require.True(t, errors.As(err, &invalidIdErr))
	require.Equal(t, domain.NewUserInvalidIdError("").Error(), err.Error())
	require.True(t, strings.HasPrefix(err.Error(), domain.USER_INVALID_ID_PREFIX))
	require.Error(t, err)
	require.Empty(t, user)
}

func TestCreateUserWithEmptyUsername(t *testing.T) {
	user, err := domain.NewUser(VALID_UUIDV4, "")

	var invalidErr *domain.UserInvalidUsernameError
	require.True(t, errors.As(err, &invalidErr))
	require.Equal(t, domain.NewUserInvalidUsernameError("").Error(), err.Error())
	require.True(t, strings.HasPrefix(err.Error(), domain.USER_INVALID_USERNAME_PREFIX))
	require.Error(t, err)
	require.Empty(t, user)
}
