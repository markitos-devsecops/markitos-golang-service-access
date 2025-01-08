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

// --------------------------------------------------------------
const USER_INVALID_NAME = "invalid name"

type UserInvalidNameError struct {
	name string
}

func NewUserInvalidNameError(name string) error {
	return &UserInvalidNameError{name: name}
}

func (e *UserInvalidNameError) Error() string {
	return fmt.Sprintf("%s: %s", USER_INVALID_NAME, e.name)
}

// --------------------------------------------------------------
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

//--------------------------------------------------------------
