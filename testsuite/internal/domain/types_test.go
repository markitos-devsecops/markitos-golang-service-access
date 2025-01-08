package domain_test

import (
	"errors"
	"strings"
	"testing"

	"markitos-golang-service-access/internal/domain"

	"github.com/stretchr/testify/require"
)

func TestCreateUserIdWithEmptyValue(t *testing.T) {
	userId, err := domain.NewUserId("")

	var invalidIdErr *domain.UserInvalidIdError
	require.True(t, errors.As(err, &invalidIdErr))
	require.Equal(t, domain.NewUserInvalidIdError("").Error(), err.Error())
	require.True(t, strings.HasPrefix(err.Error(), domain.USER_INVALID_ID))
	require.Error(t, err)
	require.Nil(t, userId)
}

func TestCreateUserIdWithInvalidFormat(t *testing.T) {
	invalidUUID := "not-a-valid-uuid"
	userId, err := domain.NewUserId(invalidUUID)

	var invalidFormatErr *domain.UserInvalidIdFormatError
	require.True(t, errors.As(err, &invalidFormatErr))
	require.Equal(t, domain.NewUserInvalidIdFormatError(invalidUUID).Error(), err.Error())
	require.True(t, strings.HasPrefix(err.Error(), domain.USER_INVALID_ID_FORMAT))
	require.Error(t, err)
	require.Nil(t, userId)
}

func TestCreateUserIdWithValidUUID(t *testing.T) {
	validUUID := "550e8400-e29b-41d4-a716-446655440000"
	userId, err := domain.NewUserId(validUUID)

	require.NoError(t, err)
	require.NotNil(t, userId)
	require.Equal(t, validUUID, userId.Value())
}

func TestCreateUserNameWithEmptyValue(t *testing.T) {
	userName, err := domain.NewUserName("")

	var invalidNameErr *domain.UserInvalidNameError
	require.True(t, errors.As(err, &invalidNameErr))
	require.Equal(t, domain.NewUserInvalidNameError("").Error(), err.Error())
	require.True(t, strings.HasPrefix(err.Error(), domain.USER_INVALID_NAME))
	require.Error(t, err)
	require.Nil(t, userName)
}

func TestCreateUserNameWithValidValue(t *testing.T) {
	validName := "John Doe"
	userName, err := domain.NewUserName(validName)

	require.NoError(t, err)
	require.NotNil(t, userName)
	require.Equal(t, validName, userName.Value())
}
