package libs

import (
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/domain/libs"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type TokenerJWT struct {
	secretKey string
}

const (
	TOKENER_JWT_TAG_FOR_EMPTY_ERROR      = "bcrypt-empty"
	TOKENER_JWT_TAG_FOR_UNEXPECTED_ERROR = "bcrypt-unexpected"
)

func NewTokenerJWT(secretKey string) libs.Tokener {
	return TokenerJWT{
		secretKey,
	}
}

func (h TokenerJWT) Create(payload string, expireAt time.Duration) (string, error) {
	if len(payload) == 0 {
		return "", domain.NewEmptyInputError(TOKENER_JWT_TAG_FOR_EMPTY_ERROR)
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(payload), 10)
	if err != nil {
		return "", domain.NewUnexpectedResultError(TOKENER_JWT_TAG_FOR_UNEXPECTED_ERROR)
	}

	return string(hashed), nil
}

func (h TokenerJWT) Validate(token string) (string, error) {
	return "", nil
}
