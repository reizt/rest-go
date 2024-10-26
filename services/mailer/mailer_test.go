package mailer

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestSend(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	apiKey := os.Getenv("SENDGRID_API_KEY")
	from := "reizt.dev@gmail.com"
	os.Setenv("SENDGRID_API_KEY", apiKey)
	os.Setenv("MAILER_FROM", from)
	s, err := New()
	require.NoError(t, err)

	// Arrange
	input := sendInput{
		To:      "reizt.dev@gmail.com",
		Subject: "Hello",
		Text:    "Hello, world!",
		Html:    "<h1>Hello, world!</h1>",
	}

	// Act
	err = s.send(input)

	// Assert
	require.NoError(t, err)
}
