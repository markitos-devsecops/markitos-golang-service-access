package domain

import (
	"fmt"
)

// --------------------------------------------------------------
const USER_INVALID_ID_PREFIX = "invalid id"

type UserInvalidIdError struct {
	id string
}

func NewUserInvalidIdError(id string) error {
	return &UserInvalidIdError{id: id}
}

func (e *UserInvalidIdError) Error() string {
	return fmt.Sprintf("%s: %s", USER_INVALID_ID_PREFIX, e.id)
}

// --------------------------------------------------------------
const USER_INVALID_USERNAME_PREFIX = "invalid username"

type UserInvalidUsernameError struct {
	username string
}

func NewUserInvalidUsernameError(username string) error {
	return &UserInvalidUsernameError{username: username}
}

func (e *UserInvalidUsernameError) Error() string {
	return fmt.Sprintf("%s: %s", USER_INVALID_USERNAME_PREFIX, e.username)
}

// --------------------------------------------------------------
const USER_INVALID_ID_FORMAT_PREFIX = "invalid id format, must be an UUIDv4"

type UserInvalidIdFormatError struct {
	id string
}

func NewUserInvalidIdFormatError(id string) error {
	return &UserInvalidIdFormatError{id: id}
}

func (e *UserInvalidIdFormatError) Error() string {
	return fmt.Sprintf("%s: %s", USER_INVALID_ID_FORMAT_PREFIX, e.id)
}

//--------------------------------------------------------------
