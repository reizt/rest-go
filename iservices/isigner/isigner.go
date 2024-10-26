package isigner

import "time"

type Service interface {
	Sign(json string, expiresIn time.Duration) (string, error)
	Verify(token string) (string, error)
}
