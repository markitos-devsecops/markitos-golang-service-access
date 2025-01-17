package dependencies_test

import (
	"errors"
	"markitos-golang-service-access/infrastructure/implementations"
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/domain/dependencies"
	"testing"

	"github.com/stretchr/testify/require"
)

var hasher dependencies.Hasher = implementations.NewHasherBCrypt()

func TestCanCreateABcryptToken(t *testing.T) {
	hashedContent, err := hasher.Create("any string value")
	require.NoError(t, err)
	require.Nil(t, err)
	require.NotEqual(t, hashedContent, "any string value")
	require.NotEmpty(t, hashedContent)
	require.NotNil(t, hashedContent)

	_, err = hasher.Create("")
	var domainError *domain.EmptyInputError
	require.True(t, errors.As(err, &domainError))
	require.NotNil(t, err)
}

func TestCanValidateABcryptToken(t *testing.T) {
	hashedContent, err := hasher.Create("any string value")
	require.NoError(t, err)
	require.Nil(t, err)

	require.True(t, hasher.Validate(hashedContent, "any string value"))
	require.False(t, hasher.Validate(hashedContent, "any string valu"))
}
