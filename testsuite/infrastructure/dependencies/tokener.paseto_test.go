package dependencies_test

import (
	"errors"
	"markitos-golang-service-access/infrastructure/implementations"
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/domain/dependencies"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	USER_TEST_TOKENER_PASETO_SECRET1 string = "any-secret-key-with-32-size-sec1"
	USER_TEST_TOKENER_PASETO_SECRET2 string = "any-secret-key-with-32-size-sec2"
)

func TestTokenerPasetoCanCreateAValidJWTToken(t *testing.T) {
	var masterTokenValue string = domain.UUIDv4()
	var duration time.Duration = time.Minute
	var issuedAt time.Time = time.Now()
	var expireAt time.Time = issuedAt.Add(duration)

	tokener := CreateTokenerPaseto(t)
	token, err := tokener.Create(masterTokenValue, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	validatedPayload, err := tokener.Validate(token)
	require.NoError(t, err)
	require.NotEmpty(t, validatedPayload)
	require.Equal(t, masterTokenValue, validatedPayload.MasterValue)
	require.WithinDuration(t, issuedAt, validatedPayload.IssueddAt, time.Second)
	require.WithinDuration(t, expireAt, validatedPayload.ExpiredAt, time.Second)
}

func TestTokenerPasetoExpiredToken(t *testing.T) {
	var masterTokenValue string = domain.UUIDv4()
	var duration time.Duration = -time.Minute

	tokener := CreateTokenerPaseto(t)
	token, err := tokener.Create(masterTokenValue, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	_, err = tokener.Validate(token)
	require.Error(t, err)
	var tokenerValidationError *domain.TokenerExpiredError
	require.True(t, errors.As(err, &tokenerValidationError))
}

func TestTokenerPasetoManipulatedToken(t *testing.T) {
	var masterTokenValue string = domain.UUIDv4()
	var duration time.Duration = time.Minute

	tokener := CreateTokenerPaseto(t)
	token, err := tokener.Create(masterTokenValue, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	parts := strings.Split(token, ".")
	parts[1] = "manipulated-payload"
	manipulatedToken := strings.Join(parts, ".")

	_, err = tokener.Validate(manipulatedToken)
	require.Error(t, err)
	var tokenerValidationError *domain.TokenerValidationError
	require.True(t, errors.As(err, &tokenerValidationError))
}

func TestTokenerPasetoInvalidSecretKeySize(t *testing.T) {
	var masterTokenValue string = domain.UUIDv4()
	var duration time.Duration = time.Minute

	tokener := CreateTokenerPaseto(t)
	token, err := tokener.Create(masterTokenValue, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	invalidTokener, _ := implementations.NewTokenerJWT(USER_TEST_TOKENER_PASETO_SECRET2)
	_, err = invalidTokener.Validate(token)
	require.Error(t, err)
	var tokenerValidationError *domain.TokenerValidationError
	require.True(t, errors.As(err, &tokenerValidationError))
}

func CreateTokenerPaseto(t *testing.T) dependencies.Tokener {
	tokener, err := implementations.NewTokenerPasseto(USER_TEST_TOKENER_PASETO_SECRET1)
	require.NoError(t, err)
	require.NotNil(t, tokener)

	return tokener
}
