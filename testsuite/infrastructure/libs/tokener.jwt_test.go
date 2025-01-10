package libs_test

import (
	"markitos-golang-service-access/infrastructure/libs"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanCreateABcryptToken(t *testing.T) {
	tokener := libs.NewHasherBCrypt()
	hashed, err := tokener.Create("any string value")

	require.NoError(t, err)
	require.Nil(t, err)
	require.NotEqual(t, hashed, "any string value")
	require.NotEmpty(t, hashed)
	require.NotNil(t, hashed)
}

func TestCanValidateABcryptToken(t *testing.T) {
	bcrypt := libs.NewHasherBCrypt()
	hashed, err := bcrypt.Create("any string value")
	require.NoError(t, err)
	require.Nil(t, err)

	require.True(t, bcrypt.Validate(hashed, "any string value"))
	require.False(t, bcrypt.Validate(hashed, "any string valu"))
}
