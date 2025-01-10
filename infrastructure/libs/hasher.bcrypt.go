package libs

import "golang.org/x/crypto/bcrypt"

type HasherBCrypt struct{}

func NewHasherBCrypt() HasherBCrypt {
	return HasherBCrypt{}
}

func (h *HasherBCrypt) Create(content string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(content), 10)

	return string(hashed), err
}
func (h *HasherBCrypt) Validate(hashedContent, rawContent string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedContent), []byte(rawContent)) == nil
}
