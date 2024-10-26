package main

import (
	"github.com/joho/godotenv"
	"github.com/reizt/rest-go/router"
	"github.com/reizt/rest-go/services"
	"github.com/reizt/rest-go/usecases"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	s := services.New()
	u := usecases.New(s)
	e := router.New(u)
	e.Logger.Fatal(e.Start(":1323"))
}
