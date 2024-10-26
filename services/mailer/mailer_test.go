package mailer

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/reizt/rest-go/iservices/imailer"
	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	apiKey := os.Getenv("SENDGRID_API_KEY")
	from := "reizt.dev@gmail.com"
	s := New(apiKey, from)

	// Arrange
	input := imailer.SendInput{
		To:      "reizt.dev@gmail.com",
		Subject: "Hello",
		Text:    "Hello, world!",
		Html:    "<h1>Hello, world!</h1>",
	}

	// Act
	err = s.Send(input)

	// Assert
	assert.NoError(t, err)
}
