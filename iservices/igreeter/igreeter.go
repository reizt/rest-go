package igreeter

type Service interface {
	Greet(name string) (string, error)
}
