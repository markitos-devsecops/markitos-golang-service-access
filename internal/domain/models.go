package domain

import (
	"time"
)

type User struct {
	Id        string    `json:"id" binding:"required,uuid"`
	Message   string    `json:"message" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required,datetime" default:"now"`
	UpdatedAt time.Time `json:"updated_at" binding:"required,datetime" default:"now"`
}

func NewUser(id, message string) (*User, error) {
	anId, anIdError := NewUserId(id)
	if anIdError != nil {
		return nil, anIdError
	}

	aMessage, aMessageError := NewUserMessage(message)
	if aMessageError != nil {
		return nil, aMessageError
	}

	return &User{
		Id:        anId.value,
		Message:   aMessage.value,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
