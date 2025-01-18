package implementations

import (
	"errors"
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/domain/dependencies"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenerJWT struct {
	secretKey string
}

const (
	TOKENER_JWT_SIMETRIC_KEY_EXACT_LENGTH = 32
)

func NewTokenerJWT(secretKey string) (dependencies.Tokener, error) {
	if len(secretKey) != TOKENER_JWT_SIMETRIC_KEY_EXACT_LENGTH {
		return nil, domain.NewTokenerInvalidKeyLengthError()
	}

	return TokenerJWT{
		secretKey: secretKey,
	}, nil
}

func (t TokenerJWT) Create(entity string, expireAt time.Duration) (string, error) {
	if len(entity) == 0 {
		return "", domain.NewTokenerInvalidKeyLengthError()
	}

	payload := dependencies.NewPayload(entity, expireAt)

	claims := jwt.MapClaims{
		dependencies.TOKENER_MASTER_VALUE_JWT_KEY: payload.Entity,
		dependencies.TOKENER_ISSUED_AT_JWT_KEY:    payload.IssuedAt.Unix(),
		dependencies.TOKENER_EXPIRED_AT_JWT_KEY:   payload.ExpiredAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(t.secretKey))
}

func (t TokenerJWT) Validate(tokenInput string) (*dependencies.Payload, error) {

	parsedToken, err := jwt.ParseWithClaims(tokenInput, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.NewTokenerValidationError("invalid signature")
		}

		return []byte(t.secretKey), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, domain.NewTokenerExpiredError()
		}
		return nil, domain.NewTokenerValidationError(err.Error())
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, domain.NewTokenerValidationError("invalid claims")
	}

	payload, err := dependencies.NewPayloadFromToken(claims)
	if err != nil {
		return nil, domain.NewTokenerValidationError("invalid token, payload")
	}

	if err := payload.Valid(); err != nil {
		return nil, err
	}

	return payload, nil
}
