package ihasher

type Service interface {
	Hash(value string) (string, error)
	Validate(value, hash string) error
}
