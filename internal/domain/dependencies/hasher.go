package dependencies

type Hasher interface {
	Create(content string) (string, error)
	Validate(hashedContent, rawContent string) bool
}
