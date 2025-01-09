package domain

import (
	"fmt"
)

// --------------------------------------------------------------
const USER_INVALID_ID = "invalid id"

type UserInvalidIdError struct {
	id string
}

func NewUserInvalidIdError(id string) error {
	return &UserInvalidIdError{id: id}
}

func (e *UserInvalidIdError) Error() string {
	return fmt.Sprintf("%s: %s", USER_INVALID_ID, e.id)
}

const USER_INVALID_ID_FORMAT = "invalid id format, must be an UUIDv4"

type UserInvalidIdFormatError struct {
	id string
}

func NewUserInvalidIdFormatError(id string) error {
	return &UserInvalidIdFormatError{id: id}
}

func (e *UserInvalidIdFormatError) Error() string {
	return fmt.Sprintf("%s: %s", USER_INVALID_ID_FORMAT, e.id)
}

// --------------------------------------------------------------
const USER_INVALID_NAME = "the name must be between 3 and 50 characters long"

type UserInvalidNameError struct {
	name string
}

func NewUserInvalidNameError(name string) error {
	return &UserInvalidNameError{name: name}
}

func (e *UserInvalidNameError) Error() string {
	return fmt.Sprintf("%s: %s", USER_INVALID_NAME, e.name)
}

const USER_INVALID_NAME_FORMAT = "the name can only contain letters and spaces and must start and end with a letter"

type UserInvalidNameFormatError struct {
	email string
}

func NewUserInvalidNameFormatError(email string) error {
	return &UserInvalidNameFormatError{email: email}
}

func (e *UserInvalidNameFormatError) Error() string {
	return fmt.Sprintf("%s: %s", USER_INVALID_EMAIL_FORMAT, e.email)
}

// --------------------------------------------------------------
const USER_INVALID_EMAIL = "invalid email or empty"

type UserInvalidEmailError struct {
	email string
}

func NewUserInvalidEmailError(email string) error {
	return &UserInvalidEmailError{email: email}
}

func (e *UserInvalidEmailError) Error() string {
	return fmt.Sprintf("%s: %s", USER_INVALID_EMAIL, e.email)
}

const USER_INVALID_EMAIL_FORMAT = "invalid email format, must be a valid email"

type UserInvalidEmailFormatError struct {
	email string
}

func NewUserInvalidEmailFormatError(email string) error {
	return &UserInvalidEmailFormatError{email: email}
}

func (e *UserInvalidEmailFormatError) Error() string {
	return fmt.Sprintf("%s: %s", USER_INVALID_EMAIL_FORMAT, e.email)
}

// --------------------------------------------------------------
const USER_INVALID_PASSWORD = "password must be between 6 and 25 characters"

type UserInvalidPasswordError struct {
	password string
}

func NewUserInvalidPasswordError(password string) error {
	return &UserInvalidPasswordError{password: password}
}

func (e *UserInvalidPasswordError) Error() string {
	return fmt.Sprintf("%s: %s", USER_INVALID_PASSWORD, e.password)
}

const USER_INVALID_PASSWORD_FORMAT = "password must contain at least 1 uppercase letter, 1 number, and 1 special character and min length 6 chars"

type UserInvalidPasswordFormatError struct {
	password string
}

func NewUserInvalidPasswordFormatError(password string) error {
	return &UserInvalidPasswordFormatError{password: password}
}

func (e *UserInvalidPasswordFormatError) Error() string {
	return fmt.Sprintf("%s: %s", USER_INVALID_PASSWORD_FORMAT, e.password)
}
