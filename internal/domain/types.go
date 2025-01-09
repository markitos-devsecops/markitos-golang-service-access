package domain

import (
	"net/mail"
	"regexp"
	"strings"
)

type UserId struct {
	value string
}

func NewUserId(value string) (*UserId, error) {
	if value == "" {
		return nil, NewUserInvalidIdError(value)
	}
	if !IsUUIDv4(value) {
		return nil, NewUserInvalidIdFormatError(value)
	}

	return &UserId{value: value}, nil
}

func (id *UserId) Value() string {
	return id.value
}

// ------------------------------------------------------------------
type UserName struct {
	value string
}

func NewUserName(value string) (*UserName, error) {
	value = strings.TrimSpace(value)
	value = regexp.MustCompile(`\s+`).ReplaceAllString(value, " ")

	if len(value) < 3 || len(value) > 50 {
		return nil, NewUserInvalidNameError(value)
	}

	regex := regexp.MustCompile(`^[a-zA-Z][a-zA-Z\s]*[a-zA-Z]$`)
	if !regex.MatchString(value) {
		return nil, NewUserInvalidNameFormatError(value)
	}

	return &UserName{value: value}, nil
}

func (msg *UserName) Value() string {
	return msg.value
}

// ---------------------------------------------------------
type UserEmail struct {
	value string
}

// RFC 5322
func NewUserEmail(value string) (*UserEmail, error) {
	if value == "" {
		return nil, NewUserInvalidEmailError(value)
	}

	if _, err := mail.ParseAddress(value); err != nil {
		return nil, NewUserInvalidEmailFormatError(value)
	}

	return &UserEmail{value: value}, nil
}

func (email *UserEmail) Value() string {
	return email.value
}

// ---------------------------------------------------------
type UserPassword struct {
	value string
}

func NewUserPassword(password string) (*UserPassword, error) {
	if password == "" {
		return nil, &UserInvalidPasswordError{}
	}

	if len(password) < 6 || len(password) > 25 {
		return nil, &UserInvalidPasswordError{}
	}

	if !isValidPassword(password) {
		return nil, &UserInvalidPasswordFormatError{}
	}

	return &UserPassword{value: password}, nil
}

func isValidPassword(password string) bool {
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecialChar := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password)

	return hasUppercase && hasNumber && hasSpecialChar
}

func (p *UserPassword) Value() string {
	return p.value
}
