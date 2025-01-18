package dependencies

import (
	"fmt"
	"markitos-golang-service-access/internal/domain"
	"time"
)

const (
	TOKENER_MASTER_VALUE_JWT_KEY = "iss"
	TOKENER_EXPIRED_AT_JWT_KEY   = "exp"
	TOKENER_ISSUED_AT_JWT_KEY    = "iat"
)

type Tokener interface {
	Create(entity string, expireAt time.Duration) (string, error)
	Validate(token string) (*Payload, error)
}

type Payload struct {
	Entity    string    `json:"master_value"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(entity string, duration time.Duration) *Payload {
	return &Payload{
		Entity:    entity,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
}

func (p *Payload) User() string {
	return p.Entity
}

func NewPayloadFromToken(parsedToken map[string]interface{}) (*Payload, error) {
	expUnix, ok := parsedToken[TOKENER_EXPIRED_AT_JWT_KEY].(float64)
	if !ok {
		return nil, domain.NewTokenerValidationError(fmt.Sprintf("error, field %s not found or with incorrect type", TOKENER_EXPIRED_AT_JWT_KEY))
	}
	expiredAt := time.Unix(int64(expUnix), 0)

	iatUnix, ok := parsedToken[TOKENER_ISSUED_AT_JWT_KEY].(float64)
	if !ok {
		return nil, domain.NewTokenerValidationError(fmt.Sprintf("error, field %s not found or with incorrect type", TOKENER_ISSUED_AT_JWT_KEY))
	}
	issuedAt := time.Unix(int64(iatUnix), 0)

	entity, ok := parsedToken[TOKENER_MASTER_VALUE_JWT_KEY].(string)
	if !ok {
		return nil, domain.NewTokenerValidationError(fmt.Sprintf("error, field %s not found or with incorrect type", TOKENER_MASTER_VALUE_JWT_KEY))
	}

	return &Payload{
		Entity:    entity,
		IssuedAt:  issuedAt,
		ExpiredAt: expiredAt,
	}, nil
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return domain.NewTokenerExpiredError()
	}

	return nil
}
