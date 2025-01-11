package libs

import (
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/domain/libs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenerJWT struct {
	secretKey string
}

const (
	TOKENER_JWT_TAG_FOR_EMPTY_ERROR = "bcrypt-empty"
	//TODO: make a constructor but dont use tag, use reason
	TOKENER_JWT_TAG_FOR_UNEXPECTED_ERROR  = "bcrypt-unexpected"
	TOKENER_JWT_SIMETRIC_KEY_EXACT_LENGTH = 32
)

func NewTokenerJWT(secretKey string) libs.Tokener {
	if len(secretKey) != TOKENER_JWT_SIMETRIC_KEY_EXACT_LENGTH {

	}
	return TokenerJWT{
		secretKey: secretKey,
	}
}

func (t TokenerJWT) Create(masterValue string, expireAt time.Duration) (string, error) {
	if len(masterValue) == 0 {
		return "", domain.NewTokenerInvalidKeyLengthError()
	}

	payload := libs.NewPayload(masterValue, expireAt)

	claims := jwt.MapClaims{
		"iss": payload.MasterValue,
		"iat": payload.IssueddAt.Unix(),
		"exp": payload.ExpiredAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(t.secretKey))
}

func (t TokenerJWT) Validate(tokenInput string) (*libs.Payload, error) {

	parsedToken, err := jwt.ParseWithClaims(tokenInput, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.NewTokenerValidationError("invalid signature")
		}

		return []byte(t.secretKey), nil
	})
	if err != nil {
		return nil, domain.NewTokenerValidationError("ivalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, domain.NewTokenerValidationError("invalid token")
	}

	payload, err := libs.NewPayloadFromToken(claims)
	if err != nil {
		return nil, domain.NewTokenerValidationError("invalid token, payload")
	}

	if err := payload.Valid(); err != nil {
		return nil, domain.NewTokenerValidationError("invalid token, payload validation")
	}

	return payload, nil
}
