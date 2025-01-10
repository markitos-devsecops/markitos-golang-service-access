package libs_test

import (
	"log"
	infraLibs "markitos-golang-service-access/infrastructure/libs"
	"markitos-golang-service-access/internal/domain"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func xTestCanCreateAJWTTokenWithAValidSecretKeyLength(t *testing.T) {
	var payload string = domain.UUIDv4()
	var duration time.Duration = time.Minute
	var issuedAt time.Time = time.Now()
	var expireAt time.Time = issuedAt.Add(duration)
	log.Println("expireAt", expireAt)

	tokener := infraLibs.NewTokenerJWT("any-secret-key-this-is-for-developer-use")
	token, err := tokener.Create(payload, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	parts := strings.Split(token, ".")
	require.Equal(t, 3, len(parts))

	validatedPayload, err := tokener.Validate(token)
	require.NoError(t, err)
	require.NotEmpty(t, validatedPayload)
	require.Equal(t, payload, validatedPayload)

}
