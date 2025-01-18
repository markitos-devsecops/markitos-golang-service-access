package services_test

import (
	"markitos-golang-service-access/internal/domain/dependencies"
	"time"
)

type MockSpyUserTokener struct {
	LastCreatedTokenMasteValue string
	LastCreatedTokenExpireAt   time.Duration
	LastValidateToken          string
}

func NewMockSpyUserTokener() *MockSpyUserTokener {
	return &MockSpyUserTokener{
		LastCreatedTokenMasteValue: "",
		LastCreatedTokenExpireAt:   0,
		LastValidateToken:          "",
	}
}

func (m *MockSpyUserTokener) Create(entity string, expireAt time.Duration) (string, error) {
	m.LastCreatedTokenMasteValue = entity
	m.LastCreatedTokenExpireAt = expireAt

	return entity, nil

}

func (m *MockSpyUserTokener) CreateHaveBeenCalledWith(entity string, expireAt time.Duration) bool {
	var result bool = m.LastCreatedTokenMasteValue == entity && m.LastCreatedTokenExpireAt == expireAt

	m.LastCreatedTokenMasteValue = ""
	m.LastCreatedTokenExpireAt = 0

	return result
}

func (m *MockSpyUserTokener) Validate(token string) (*dependencies.Payload, error) {
	m.LastValidateToken = token

	return nil, nil
}

func (m *MockSpyUserTokener) ValidateHaveBeenCalledWith(token string) bool {
	var result bool = m.LastValidateToken == token

	m.LastValidateToken = ""

	return result
}
