package idatabase

type User struct {
	Id           string
	Name         string
	Email        string
	PasswordHash string
}

type UserUpdate struct {
	Name         string
	Email        string
	PasswordHash string
}

type UserRepo interface {
	GetById(id string) (User, error)
	GetByEmail(email string) (User, error)
	Create(data User) error
	Update(id string, data UserUpdate) error
	Delete(id string) error
}
