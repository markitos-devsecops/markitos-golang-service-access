package domain

type Hasher interface {
	Create(content string) (string, error)
	Validate(content string) error
}
