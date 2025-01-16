package api_test

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

func (m *MockSpyUserTokener) Create(masterValue string, expireAt time.Duration) (string, error) {
	m.LastCreatedTokenMasteValue = masterValue
	m.LastCreatedTokenExpireAt = expireAt

	return masterValue, nil

}

func (m *MockSpyUserTokener) CreateHaveBeenCalledWith(masterValue string, expireAt time.Duration) bool {
	var result bool = m.LastCreatedTokenMasteValue == masterValue && m.LastCreatedTokenExpireAt == expireAt

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
