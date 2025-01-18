package dependencies_test

import (
	"errors"
	"markitos-golang-service-access/infrastructure/implementations"
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/domain/dependencies"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"
)

const (
	USER_TEST_TOKENER_JWT_SECRET1 string = "any-secret-key-with-32-size-sec1"
	USER_TEST_TOKENER_JWT_SECRET2 string = "any-secret-key-with-32-size-sec2"
)

func TestTokenerJWTCanCreateAValidJWTToken(t *testing.T) {
	var entity string = domain.UUIDv4()
	var duration time.Duration = time.Minute
	var issuedAt time.Time = time.Now()
	var expireAt time.Time = issuedAt.Add(duration)

	tokener := CreateTokenerJWT(t)
	token, err := tokener.Create(entity, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	parts := strings.Split(token, ".")
	require.Equal(t, 3, len(parts))

	validatedPayload, err := tokener.Validate(token)
	require.NoError(t, err)
	require.NotEmpty(t, validatedPayload)
	require.Equal(t, entity, validatedPayload.Entity)
	require.WithinDuration(t, issuedAt, validatedPayload.IssuedAt, time.Second)
	require.WithinDuration(t, expireAt, validatedPayload.ExpiredAt, time.Second)
}

func TestTokenerJWTExpiredToken(t *testing.T) {
	var entity string = domain.UUIDv4()
	var duration time.Duration = -time.Minute

	tokener := CreateTokenerJWT(t)
	token, err := tokener.Create(entity, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	_, err = tokener.Validate(token)
	require.Error(t, err)
	var tokenerValidationError *domain.TokenerExpiredError
	require.True(t, errors.As(err, &tokenerValidationError))
}

func TestTokenerJWTInvalidSignatureToken(t *testing.T) {
	var entity string = domain.UUIDv4()
	var duration time.Duration = time.Minute

	tokener := CreateTokenerJWT(t)
	token, err := tokener.Create(entity, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	parts := strings.Split(token, ".")
	parts[2] = "invalid-signature"
	invalidToken := strings.Join(parts, ".")

	_, err = tokener.Validate(invalidToken)
	require.Error(t, err)
	var tokenerValidationError *domain.TokenerValidationError
	require.True(t, errors.As(err, &tokenerValidationError))
}

func TestTokenerJWTManipulatedToken(t *testing.T) {
	var entity string = domain.UUIDv4()
	var duration time.Duration = time.Minute

	tokener := CreateTokenerJWT(t)
	token, err := tokener.Create(entity, duration)
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

func TestTokenerJWTInvalidSecretKeySize(t *testing.T) {
	var entity string = domain.UUIDv4()
	var duration time.Duration = time.Minute

	tokener := CreateTokenerJWT(t)
	token, err := tokener.Create(entity, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	invalidTokener, _ := implementations.NewTokenerJWT(USER_TEST_TOKENER_JWT_SECRET2)
	_, err = invalidTokener.Validate(token)
	require.Error(t, err)
	var tokenerValidationError *domain.TokenerValidationError
	require.True(t, errors.As(err, &tokenerValidationError))
}

func TestTokenerJWTNoneAlgorithmAttack(t *testing.T) {
	var entity string = domain.UUIDv4()
	var duration time.Duration = time.Minute

	tokener := CreateTokenerJWT(t)
	token, err := tokener.Create(entity, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	parts := strings.Split(token, ".")
	header := `{"alg":"none","typ":"JWT"}`
	parts[0] = jwt.EncodeSegment([]byte(header))
	noneToken := strings.Join(parts, ".")

	_, err = tokener.Validate(noneToken)
	require.Error(t, err)
	var tokenerValidationError *domain.TokenerValidationError
	require.True(t, errors.As(err, &tokenerValidationError))
}

func CreateTokenerJWT(t *testing.T) dependencies.Tokener {
	tokener, err := implementations.NewTokenerJWT(USER_TEST_TOKENER_JWT_SECRET1)
	require.NoError(t, err)
	require.NotNil(t, tokener)

	return tokener
}
