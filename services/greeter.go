package services

import (
	"fmt"
)

type Greeter struct{}

func (g Greeter) Greet(name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}
