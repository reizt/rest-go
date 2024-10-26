package greeter

import (
	"fmt"

	"github.com/reizt/rest-go/iservices/igreeter"
)

func New() igreeter.Service {
	return Service{}
}

type Service struct{}

func (g Service) Greet(name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}
