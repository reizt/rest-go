package iservices

type Greeter interface {
	Greet(name string) (string, error)
}
