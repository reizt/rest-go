package database

import (
	"context"
	"fmt"
	"os"
)

func Clean() {
	development := os.Getenv("TEST_CLEAR_DATABASE")
	if development != "on" {
		fmt.Println("🚫 Not in development mode")
		return
	}
	client, err := getClient()
	if err != nil {
		panic(err)
	}
	userRepo := UserRepo{client}
	codeRepo := CodeRepo{client}

	ctx := context.Background()
	userRepo.deleteAll(ctx)
	codeRepo.deleteAll(ctx)

	fmt.Println("🎉 Database cleared")
}
