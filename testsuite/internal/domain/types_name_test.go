package domain_test

import (
	"errors"
	"strings"
	"testing"

	"markitos-golang-service-access/internal/domain"

	"github.com/stretchr/testify/require"
)

func TestCreateUserNameWithEmptyValue(t *testing.T) {
	userName, err := domain.NewUserName("")

	var invalidNameErr *domain.UserInvalidNameError
	require.True(t, errors.As(err, &invalidNameErr))
	require.Error(t, err)
	require.Nil(t, userName)
}

func TestCreateUserNameWithTooShortValue(t *testing.T) {
	userName, err := domain.NewUserName("Jo")

	var invalidNameErr *domain.UserInvalidNameError
	require.True(t, errors.As(err, &invalidNameErr))
	require.Error(t, err)
	require.Nil(t, userName)
}

func TestCreateUserNameWithTooLongValue(t *testing.T) {
	longName := strings.Repeat("a", 51)
	userName, err := domain.NewUserName(longName)

	var invalidNameErr *domain.UserInvalidNameError
	require.True(t, errors.As(err, &invalidNameErr))
	require.Error(t, err)
	require.Nil(t, userName)
}

func TestCreateUserNameWithInvalidCharacters(t *testing.T) {
	invalidName := "John_Doe!"
	userName, err := domain.NewUserName(invalidName)

	var formatError *domain.UserInvalidNameFormatError
	require.True(t, errors.As(err, &formatError))
	require.Error(t, err)
	require.Nil(t, userName)
}

func TestCreateUserNameWithMultipleSpaces(t *testing.T) {
	validName := "John     Doe"
	expectedName := "John Doe"
	userName, err := domain.NewUserName(validName)

	require.NoError(t, err)
	require.NotNil(t, userName)
	require.Equal(t, expectedName, userName.Value())
}

func TestCreateUserNameWithValidValue(t *testing.T) {
	validName := "John Doe"
	userName, err := domain.NewUserName(validName)

	require.NoError(t, err)
	require.NotNil(t, userName)
	require.Equal(t, validName, userName.Value())
}
