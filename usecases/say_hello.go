package usecases

import (
	"fmt"

	"github.com/reizt/rest-go/iservices"
	"github.com/reizt/rest-go/iusecases"
)

func SayHello(s *iservices.All) iusecases.SayHello {
	return func(input iusecases.SayHelloInput) (iusecases.SayHelloOutput, error) {
		output := iusecases.SayHelloOutput{
			Message: fmt.Sprintf("Hello, %s!", input.Name),
		}
		return output, nil
	}
}
