package mservices

type Hasher struct {
	Hash_     func(value string) (string, error)
	Validate_ func(value, hash string) error
}

func (h Hasher) Hash(value string) (string, error) {
	return h.Hash_(value)
}

func (h Hasher) Validate(value, hash string) error {
	return h.Validate_(value, hash)
}
