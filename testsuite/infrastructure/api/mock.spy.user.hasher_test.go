package api_test

type MockSpyUserHasher struct {
	LastCreatedHash  string
	LastValidateHash string
}

func NewMockSpyUserHasher() *MockSpyUserHasher {
	return &MockSpyUserHasher{
		LastCreatedHash:  "",
		LastValidateHash: "",
	}
}

func (m *MockSpyUserHasher) Create(unhashed string) (string, error) {
	m.LastCreatedHash = unhashed

	m.LastCreatedHash = ""

	return unhashed, nil

}

func (m *MockSpyUserHasher) CreateHaveBeenCalledWith(unhashed string) bool {
	var result bool = m.LastCreatedHash == unhashed

	m.LastCreatedHash = ""

	return result
}

func (m *MockSpyUserHasher) Validate(hashed, unhashed string) bool {
	m.LastValidateHash = hashed

	m.LastValidateHash = ""

	return hashed == unhashed
}

func (m *MockSpyUserHasher) ValidateHaveBeenCalledWith(hashed, unhashed string) bool {
	var result bool = m.LastValidateHash == hashed

	m.LastValidateHash = ""

	return result
}
