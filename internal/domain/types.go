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

type UserMessage struct {
	value string
}

func NewUserMessage(value string) (*UserMessage, error) {
	if value == "" {
		return nil, NewUserInvalidMessageError(value)
	}
	return &UserMessage{value: value}, nil
}

func (msg *UserMessage) Value() string {
	return msg.value
}
