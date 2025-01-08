package domain_test

import (
	"markitos-golang-service-access/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUserId(t *testing.T) {
	t.Run("should create UserId when value is a valid UUIDv4", func(t *testing.T) {
		validUUID := "550e8400-e29b-41d4-a716-446655440000"
		userId, err := domain.NewUserId(validUUID)
		require.NoError(t, err)
		assert.NotNil(t, userId)
		assert.Equal(t, validUUID, userId.Value())
	})

	t.Run("should fail when value is empty", func(t *testing.T) {
		emptyValue := ""
		userId, err := domain.NewUserId(emptyValue)
		require.Error(t, err)
		assert.Nil(t, userId)
		assert.EqualError(t, err, domain.NewUserInvalidIdError(emptyValue).Error())
	})

	t.Run("should fail when value is not a valid UUIDv4", func(t *testing.T) {
		invalidUUID := "not-a-valid-uuid"
		userId, err := domain.NewUserId(invalidUUID)
		require.Error(t, err)
		assert.Nil(t, userId)
		assert.EqualError(t, err, domain.NewUserInvalidIdFormatError(invalidUUID).Error())
	})
}

func TestNewUserName(t *testing.T) {
	t.Run("should create UserName when value is valid", func(t *testing.T) {
		validName := "John Doe"
		userName, err := domain.NewUserName(validName)
		require.NoError(t, err)
		assert.NotNil(t, userName)
		assert.Equal(t, validName, userName.Value())
	})

	t.Run("should fail when value is empty", func(t *testing.T) {
		emptyValue := ""
		userName, err := domain.NewUserName(emptyValue)
		require.Error(t, err)
		assert.Nil(t, userName)
		assert.EqualError(t, err, domain.NewUserInvalidNameError(emptyValue).Error())
	})
}
