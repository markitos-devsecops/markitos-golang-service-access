package domain

import (
	"errors"
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
	value string
}

func NewUserInvalidNameFormatError(value string) error {
	return &UserInvalidNameFormatError{value: value}
}

func (e *UserInvalidNameFormatError) Error() string {
	return fmt.Sprintf("%s: %s", USER_INVALID_NAME_FORMAT, e.value)
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

// --------------------------------------------------------------

const ERROR_EMPTY_INPUT = "cant pass an empty value"

type EmptyInputError struct {
	tag string
}

func NewEmptyInputError(tag string) error {
	return &EmptyInputError{tag}
}

func (e *EmptyInputError) Error() string {
	return fmt.Sprintf("%s: %s", ERROR_EMPTY_INPUT, e.tag)
}

// --------------------------------------------------------------

const ERROR_UNEXPECTED_RESULT = "unexpected result"

type UnexpectedResultError struct {
	tag string
}

func NewUnexpectedResultError(tag string) error {
	return &UnexpectedResultError{tag}
}

func (e *UnexpectedResultError) Error() string {
	return fmt.Sprintf("%s: %s", ERROR_UNEXPECTED_RESULT, e.tag)
}

// --------------------------------------------------------------

const ERROR_TOKENER_INVALID_KEY_SIZE = "secret key must contain exactly 32 chars"

type TokenerInvalidKeyLengthError struct {
}

func NewTokenerInvalidKeyLengthError() error {
	return &TokenerInvalidKeyLengthError{}
}

func (e *TokenerInvalidKeyLengthError) Error() string {
	return fmt.Sprint(ERROR_TOKENER_INVALID_KEY_SIZE)
}

const ERROR_TOKENER_CREATION = "cant create a new token"

type TokenerCreationError struct {
}

func NewTokenerCreationError() error {
	return &TokenerCreationError{}
}

func (e *TokenerCreationError) Error() string {
	return fmt.Sprint(ERROR_TOKENER_CREATION)
}

const ERROR_TOKENER_VALIDATION = "cant validate token, reason: %s"

type TokenerValidationError struct {
	reason string
}

func NewTokenerValidationError(reason string) error {
	return &TokenerValidationError{
		reason: reason,
	}
}

func (e *TokenerValidationError) Error() string {
	return fmt.Sprintf(ERROR_TOKENER_VALIDATION, e.reason)
}

const ERROR_TOKENER_EXPIRED = "token have been expired"

type TokenerExpiredError struct {
}

func NewTokenerExpiredError() error {
	return &TokenerExpiredError{}
}

func (e *TokenerExpiredError) Error() string {
	return ERROR_TOKENER_VALIDATION
}

const ERROR_USER_UNAUTHORIZED = "authorized action"

type UserUnauthorizedError struct {
}

func NewUnauthorizedError() error {
	return &UserUnauthorizedError{}
}

func (e *UserUnauthorizedError) Error() string {
	return ERROR_USER_UNAUTHORIZED
}

var (
	ErrUserAlreadyExists error = errors.New("user already exists")
)
