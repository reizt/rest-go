package services

import "github.com/reizt/rest-go/iservices"

func New() *iservices.All {
	return &iservices.All{
		Greeter: Greeter{},
	}
}
