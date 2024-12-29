package domain

import (
	"time"
)

type User struct {
	Id        string    `json:"id" binding:"required,uuid"`
	Username  string    `json:"username" binding:"required,email"`
	CreatedAt time.Time `json:"created_at" binding:"required,datetime" default:"now"`
	UpdatedAt time.Time `json:"updated_at" binding:"required,datetime" default:"now"`
}

func NewUser(id, username string) (*User, error) {
	anId, anIdError := NewUserId(id)
	if anIdError != nil {
		return nil, anIdError
	}

	aUsername, aUsernameError := NewUsername(username)
	if aUsernameError != nil {
		return nil, aUsernameError
	}

	return &User{
		Id:        anId.value,
		Username:  aUsername.value,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
