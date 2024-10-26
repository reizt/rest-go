package idatabase

type Code struct {
	Id        string
	Email     string
	Action    string
	ValueHash string
	ExpiresAt int64
	CreatedAt int64
}

type CodeRepo interface {
	GetById(id string) (Code, error)
	GetByEmail(email string) (Code, error)
	Create(data Code) error
	Delete(id string) error
}
