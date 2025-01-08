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

type UserName struct {
	value string
}

func NewUserName(value string) (*UserName, error) {
	if value == "" {
		return nil, NewUserInvalidNameError(value)
	}
	return &UserName{value: value}, nil
}

func (msg *UserName) Value() string {
	return msg.value
}
