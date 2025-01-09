package domain_test

import (
	"testing"

	"markitos-golang-service-access/internal/domain"

	"github.com/stretchr/testify/require"
)

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
		{"invalidpasswordnoupper"},
		{"NoNumberOrSpecialChar"},
		{"NoSpecialChar12345"},
		{"1234567@noupper"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			_, err := domain.NewUserPassword(tt.password)
			require.Error(t, err)
			require.IsType(t, &domain.UserInvalidPasswordFormatError{}, err)
		})
	}
}
