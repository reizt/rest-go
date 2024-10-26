package database

import (
	"context"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/reizt/rest-go/ent"
	"github.com/reizt/rest-go/iservices/idatabase"
)

func New() idatabase.Service {
	addr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	client, err := ent.Open("postgres", addr)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Schema.Create(ctx); err != nil {
		panic(err)
	}
	return idatabase.Service{
		User: UserRepo{client},
		Code: CodeRepo{client},
	}
}
