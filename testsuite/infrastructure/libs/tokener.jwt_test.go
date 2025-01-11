package libs_test

import (
	"errors"
	"fmt"
	"markitos-golang-service-access/infrastructure/libs"
	"markitos-golang-service-access/internal/domain"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"
)

const (
	USER_TEST_TOKENER_JWT_SECRET string = "any-secret-key-this-is-for-developer-use-32"
)

func TestCanCreateAValidJWTToken(t *testing.T) {
	var masterTokenValue string = domain.UUIDv4()
	var duration time.Duration = time.Minute
	var issuedAt time.Time = time.Now()
	var expireAt time.Time = issuedAt.Add(duration)

	tokener := libs.NewTokenerJWT(USER_TEST_TOKENER_JWT_SECRET)
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

func TestExpiredToken(t *testing.T) {
	var masterTokenValue string = domain.UUIDv4()
	var duration time.Duration = -time.Minute

	tokener := libs.NewTokenerJWT(USER_TEST_TOKENER_JWT_SECRET)
	token, err := tokener.Create(masterTokenValue, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	_, err = tokener.Validate(token)
	fmt.Println("Error: ", err)
	require.Error(t, err)
	var tokenerValidationError *domain.TokenerExpiredError
	require.True(t, errors.As(err, &tokenerValidationError))
}

func TestInvalidSignatureToken(t *testing.T) {
	var masterTokenValue string = domain.UUIDv4()
	var duration time.Duration = time.Minute

	tokener := libs.NewTokenerJWT(USER_TEST_TOKENER_JWT_SECRET)
	token, err := tokener.Create(masterTokenValue, duration)
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

func TestManipulatedToken(t *testing.T) {
	var masterTokenValue string = domain.UUIDv4()
	var duration time.Duration = time.Minute

	tokener := libs.NewTokenerJWT(USER_TEST_TOKENER_JWT_SECRET)
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

func TestInvalidSecretKey(t *testing.T) {
	var masterTokenValue string = domain.UUIDv4()
	var duration time.Duration = time.Minute

	tokener := libs.NewTokenerJWT(USER_TEST_TOKENER_JWT_SECRET)
	token, err := tokener.Create(masterTokenValue, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	invalidTokener := libs.NewTokenerJWT("different-secret-key")
	_, err = invalidTokener.Validate(token)
	require.Error(t, err)
	var tokenerValidationError *domain.TokenerValidationError
	require.True(t, errors.As(err, &tokenerValidationError))
}

func TestNoneAlgorithmAttack(t *testing.T) {
	var masterTokenValue string = domain.UUIDv4()
	var duration time.Duration = time.Minute

	tokener := libs.NewTokenerJWT(USER_TEST_TOKENER_JWT_SECRET)
	token, err := tokener.Create(masterTokenValue, duration)
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
