package libs_test

import (
	"markitos-golang-service-access/infrastructure/libs"
	"markitos-golang-service-access/internal/domain"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCanCreateAValidJWTToken(t *testing.T) {
	var masterTokenValue string = domain.UUIDv4()
	var duration time.Duration = time.Minute
	var issuedAt time.Time = time.Now()
	var expireAt time.Time = issuedAt.Add(duration)

	tokener := libs.NewTokenerJWT("any-secret-key-this-is-for-developer-use")
	token, err := tokener.Create(masterTokenValue, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	parts := strings.Split(token, ".")
	require.Equal(t, 3, len(parts))

	validatedPayload, err := tokener.Validate(token)
	require.NoError(t, err)
	require.NotEmpty(t, validatedPayload)
	require.Equal(t, masterTokenValue, validatedPayload.MasterValue)
	require.WithinDuration(t, issuedAt, validatedPayload.IssueddAt, time.Second)
	require.WithinDuration(t, expireAt, validatedPayload.ExpiredAt, time.Second)
}
