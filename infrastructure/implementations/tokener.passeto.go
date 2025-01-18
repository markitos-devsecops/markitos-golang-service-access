package implementations

import (
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/domain/dependencies"
	"time"

	"github.com/o1egl/paseto"
)

type TokenerPasseto struct {
	secretKey []byte
	paseto    paseto.V2
}

const (
	TOKENER_PASSETO_CHACHA20POLY_KEYSIZE = 32
)

func NewTokenerPasseto(secretKey string) (dependencies.Tokener, error) {
	if len(secretKey) != TOKENER_PASSETO_CHACHA20POLY_KEYSIZE {
		return nil, domain.NewTokenerInvalidKeyLengthError()
	}

	return TokenerPasseto{
		secretKey: []byte(secretKey),
		paseto:    *paseto.NewV2(),
	}, nil
}

func (t TokenerPasseto) Create(entity string, expireAt time.Duration) (string, error) {
	if len(entity) == 0 {
		return "", domain.NewTokenerInvalidKeyLengthError()
	}

	payload := dependencies.NewPayload(entity, expireAt)
	token, err := t.paseto.Encrypt(t.secretKey, payload, nil)
	if err != nil {
		return "", domain.NewTokenerValidationError("error creating token " + err.Error())
	}

	return token, nil
}

func (t TokenerPasseto) Validate(tokenInput string) (*dependencies.Payload, error) {
	payload := &dependencies.Payload{}

	err := t.paseto.Decrypt(tokenInput, t.secretKey, payload, nil)
	if err != nil {
		return nil, domain.NewTokenerValidationError("invalid token " + err.Error())
	}

	if err := payload.Valid(); err != nil {
		return nil, err
	}

	return payload, nil
}
