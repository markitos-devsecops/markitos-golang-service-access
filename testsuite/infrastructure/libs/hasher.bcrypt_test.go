package libs_test

import (
	"errors"
	"markitos-golang-service-access/infrastructure/libs"
	infraLibs "markitos-golang-service-access/infrastructure/libs"
	"markitos-golang-service-access/internal/domain"
	domainLibs "markitos-golang-service-access/internal/domain/libs"
	"testing"

	"github.com/stretchr/testify/require"
)

var hasher domainLibs.Hasher = infraLibs.NewHasherBCrypt()

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
	bcrypt := libs.NewHasherBCrypt()
	hashedContent, err := bcrypt.Create("any string value")
	require.NoError(t, err)
	require.Nil(t, err)

	require.True(t, bcrypt.Validate(hashedContent, "any string value"))
	require.False(t, bcrypt.Validate(hashedContent, "any string valu"))
}
