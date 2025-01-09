package domain_test

import (
	"errors"
	"strings"
	"testing"

	"markitos-golang-service-access/internal/domain"

	"github.com/stretchr/testify/require"
)

func TestCreateUserEmailWithEmptyValue(t *testing.T) {
	userId, err := domain.NewUserEmail("")

	var invalidErr *domain.UserInvalidEmailError
	require.True(t, errors.As(err, &invalidErr))
	require.Equal(t, domain.NewUserInvalidEmailError("").Error(), err.Error())
	require.True(t, strings.HasPrefix(err.Error(), domain.USER_INVALID_EMAIL))
	require.Error(t, err)
	require.Nil(t, userId)
}

func TestCreateUserEmailWithValidFormat(t *testing.T) {
	tests := []struct {
		name      string
		email     string
		expectErr bool
		errType   error
	}{
		{"ValidEmail: simple", "juan.perez@example.com", false, nil},
		{"ValidEmail: uppercase", "JUAN.PEREZ@EXAMPLE.COM", false, nil},
		{"ValidEmail: with number", "juan123@example.com", false, nil},
		{"ValidEmail: with hyphen", "juan-perez@example.com", false, nil},
		{"ValidEmail: with dot", "juan.perez.sanchez@example.com", false, nil},
		{"ValidEmail: with plus", "juan+alias@example.com", false, nil},
		{"ValidEmail: subdomain", "juan@subdomain.example.com", false, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := domain.NewUserEmail(tt.email)

			if tt.expectErr {
				require.Error(t, err)
				require.IsType(t, tt.errType, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestCreateUserEmailWithInvalidFormat(t *testing.T) {
	tests := []struct {
		name      string
		email     string
		expectErr bool
		errType   error
	}{
		{"InvalidEmail: missing domain", "juan@", true, &domain.UserInvalidEmailFormatError{}},
		{"InvalidEmail: missing username", "@example.com", true, &domain.UserInvalidEmailFormatError{}},
		{"InvalidEmail: double dot in domain", "juan@example..com", true, &domain.UserInvalidEmailFormatError{}},
		{"InvalidEmail: missing @ symbol", "juanexample.com", true, &domain.UserInvalidEmailFormatError{}},
		{"InvalidEmail: invalid character in domain", "juan@ex@mple.com", true, &domain.UserInvalidEmailFormatError{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := domain.NewUserEmail(tt.email)

			if tt.expectErr {
				require.Error(t, err)
				require.IsType(t, tt.errType, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

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

func TestUserPassword_Valid(t *testing.T) {
	tests := []struct {
		password string
	}{
		{"Password123@"},
		{"Valid1@Password"},
		{"Test@123"},
	}

	for _, tt := range tests {
		t.Run(tt.password, func(t *testing.T) {
			password, err := domain.NewUserPassword(tt.password)
			require.NoError(t, err)
			require.Equal(t, tt.password, password.Value())
		})
	}
}

func TestUserPassword_EmptyOrInvalidLength(t *testing.T) {
	tests := []struct {
		password string
	}{
		{""},
		{"short"},
		{"thisisaverylongpassword12345678"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			_, err := domain.NewUserPassword(tt.password)
			require.Error(t, err)

			require.IsType(t, &domain.UserInvalidPasswordError{}, err)
		})
	}
}

func TestUserPassword_InvalidFormat(t *testing.T) {
	tests := []struct {
		password string
	}{
		{"invalidpassword"},       // No contiene mayúscula, número ni carácter especial
		{"NoNumberOrSpecialChar"}, // No contiene número ni carácter especial
		{"NoSpecialChar12345"},    // No contiene carácter especial
		{"1234567@"},              // No contiene mayúscula
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			_, err := domain.NewUserPassword(tt.password)
			require.Error(t, err)
			require.IsType(t, &domain.UserInvalidPasswordFormatError{}, err)
		})
	}
}
