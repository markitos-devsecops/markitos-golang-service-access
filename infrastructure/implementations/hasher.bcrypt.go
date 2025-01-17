package implementations

import (
	"markitos-golang-service-access/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

type HasherBCrypt struct {
}

const (
	HASHER_BCRYPT_TAG_FOR_EMPTY_ERROR      = "bcrypt-empty"
	HASHER_BCRYPT_TAG_FOR_UNEXPECTED_ERROR = "bcrypt-unexpected"
)

func NewHasherBCrypt() HasherBCrypt {
	return HasherBCrypt{}
}

func (h HasherBCrypt) Create(content string) (string, error) {
	if len(content) == 0 {
		return "", domain.NewEmptyInputError(HASHER_BCRYPT_TAG_FOR_EMPTY_ERROR)
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(content), bcrypt.DefaultCost)
	if err != nil {
		return "", domain.NewUnexpectedResultError(HASHER_BCRYPT_TAG_FOR_EMPTY_ERROR)
	}

	return string(hashed), nil
}
func (h HasherBCrypt) Validate(hashedContent, rawContent string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedContent), []byte(rawContent))
	if err != nil {
		return false
	}

	return true
}
