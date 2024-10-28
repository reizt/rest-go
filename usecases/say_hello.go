package usecases

import (
	"context"
	"fmt"

	"reij.uno/iservices"
	i "reij.uno/iusecases"
)

func SayHello(s *iservices.All) i.SayHello {
	return func(input i.SayHelloInput, ctx context.Context) (*i.SayHelloOutput, error) {
		output := i.SayHelloOutput{
			Message: fmt.Sprintf("Hello, %s!", input.Name),
		}
		return &output, nil
	}
}
