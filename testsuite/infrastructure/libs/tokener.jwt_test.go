package libs_test

import (
	"markitos-golang-service-access/infrastructure/libs"
	"markitos-golang-service-access/internal/domain"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanCreateAJWTTokenWithAValidUUID(t *testing.T) {
	tokener := libs.NewTokenerJWT("secret-value")
	token, err := tokener.Create(domain.UUIDv4())

	require.NoError(t, err)
	require.NotEmpty(t, token)
	parts := strings.Split(token, ".")
	require.Equal(t, 3, len(parts))
}
