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
