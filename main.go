package main

import (
	"github.com/joho/godotenv"
	"reij.uno/router"
	"reij.uno/services"
	"reij.uno/usecases"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	s, err := services.New()
	if err != nil {
		panic(err)
	}
	u := usecases.New(s)
	e := router.New(u)
	e.Logger.Fatal(e.Start(":1323"))
}
