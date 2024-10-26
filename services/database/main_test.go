package database

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/reizt/rest-go/ent"
)

var (
	client *ent.Client
)

func cleanup() {
	userRepo := UserRepo{client}
	codeRepo := CodeRepo{client}

	ctx := context.Background()
	userRepo.deleteAll(ctx)
	codeRepo.deleteAll(ctx)
}

func TestMain(m *testing.M) {
	fmt.Println("🎯 Before all")
	godotenv.Load("../../.env")
	client, _ = getClient()

	fmt.Println("🎯 Run tests")
	code := m.Run()

	fmt.Println("🎯 After all")
	os.Exit(code)
}
