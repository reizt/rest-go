package usecases

import (
	"context"
	"fmt"

	"github.com/reizt/rest-go/iservices"
	i "github.com/reizt/rest-go/iusecases"
)

func SayHello(s *iservices.All) i.SayHello {
	return func(input i.SayHelloInput, ctx context.Context) (*i.SayHelloOutput, error) {
		output := i.SayHelloOutput{
			Message: fmt.Sprintf("Hello, %s!", input.Name),
		}
		return &output, nil
	}
}
