package libs

import "time"

type Tokener interface {
	Create(payload string, expireAt time.Duration) (string, error)
	Validate(token string) (string, error)
}
