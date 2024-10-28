package greeter

import (
	"fmt"

	"reij.uno/iservices/igreeter"
)

type service struct{}

func New() igreeter.Service {
	return service{}
}

func (g service) Greet(name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}
