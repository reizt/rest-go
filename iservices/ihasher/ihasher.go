package ihasher

type Service interface {
	Hash(password string) (string, error)
	Validate(password, hash string) error
}
