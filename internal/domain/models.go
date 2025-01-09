package domain

import (
	"time"
)

type User struct {
	Id        string    `json:"id" binding:"required,uuid"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"-" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required,datetime" default:"now"`
	UpdatedAt time.Time `json:"updated_at" binding:"required,datetime" default:"now"`
}

func NewUser(id, name, email, password string) (*User, error) {
	anId, err := NewUserId(id)
	if err != nil {
		return nil, err
	}

	aName, err := NewUserName(name)
	if err != nil {
		return nil, err
	}

	anEmail, err := NewUserEmail(email)
	if err != nil {
		return nil, err
	}

	aPassword, err := NewUserPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		Id:        anId.value,
		Name:      aName.value,
		Email:     anEmail.value,
		Password:  aPassword.value,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
