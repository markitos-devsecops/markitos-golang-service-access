package domain

type UserId struct {
	value string
}

func NewUserId(value string) (*UserId, error) {
	if value == "" {
		return nil, NewUserInvalidIdError(value)
	}
	if !IsUUIDv4(value) {
		return nil, NewUserInvalidIdFormatError(value)
	}

	return &UserId{value: value}, nil
}

func (id *UserId) Value() string {
	return id.value
}

type Username struct {
	value string
}

func NewUsername(value string) (*Username, error) {
	if value == "" {
		return nil, NewUserInvalidUsernameError(value)
	}
	return &Username{value: value}, nil
}

func (msg *Username) Value() string {
	return msg.value
}
