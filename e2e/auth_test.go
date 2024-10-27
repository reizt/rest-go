package e2e

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/reizt/rest-go/e2e/fetcher"
	"github.com/reizt/rest-go/entities"
	"github.com/reizt/rest-go/handlers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	userName     = "John"
	userNameNew  = "Jane"
	userEmail    = "reizt.dev@gmail.com"
	userPassword = "password"
)

func clearDatabase(t *testing.T) {
	req := fetcher.Request{
		Method: "POST",
		Path:   "/dev/clear-database",
		Body:   nil,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	f := fetcher.New("http://localhost:1323")
	_, err := f.Fetch(req)
	require.NoError(t, err)
}

func testIssueCode(t *testing.T, f *fetcher.Fetcher) string {
	// Arrange
	req := fetcher.Request{
		Method: "POST",
		Path:   "/auth/code/issue",
		Body: handlers.IssueCodeReqBody{
			Email:  userEmail,
			Action: entities.CodeActionCreateUser,
		},
		Headers: map[string]string{},
	}

	// Act
	resp, err := f.Fetch(req)

	// Assert
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.LessOrEqual(t, resp.StatusCode, 209)
	respBody := handlers.IssueCodeResBody{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	require.NoError(t, err)
	assert.NotEmpty(t, respBody.CodeId, "issue code response should contain code id")

	return respBody.CodeId
}

func testVerifyCode(t *testing.T, f *fetcher.Fetcher, codeId string) string {
	// Arrange
	codeValue := os.Getenv("TEST_GENERATE_CODE_FIXED_VALUE")
	if codeValue == "" {
		fmt.Println("TEST_GENERATE_CODE_FIXED_VALUE is not set")
		t.Fail()
	}
	req := fetcher.Request{
		Method: "POST",
		Path:   "/auth/code/verify",
		Body: handlers.VerifyCodeReqBody{
			CodeId: codeId,
			Code:   codeValue,
		},
		Headers: map[string]string{},
	}

	// Act
	resp, err := f.Fetch(req)

	// Assert
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.LessOrEqual(t, resp.StatusCode, 209)
	otpToken := resp.Cookies[handlers.OTPTokenCookieName]
	assert.NotEmpty(t, otpToken, "verify code response should contain token cookie")
	return otpToken
}

func testCreateUser(t *testing.T, f *fetcher.Fetcher, otpToken string) string {
	// Arrange
	req := fetcher.Request{
		Method: "POST",
		Path:   "/auth/create-user",
		Body: handlers.CreateUserReqBody{
			Name:     userName,
			Password: userPassword,
		},
		Cookies: map[string]string{
			handlers.OTPTokenCookieName: otpToken,
		},
	}

	// Act
	resp, err := f.Fetch(req)

	// Assert
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.LessOrEqual(t, resp.StatusCode, 209)
	loginToken := resp.Cookies[handlers.LoginTokenCookieName]
	assert.NotEmpty(t, loginToken, "create user response should contain login token cookie")

	return loginToken
}

func testGetUser(t *testing.T, f *fetcher.Fetcher, loginToken string) {
	// Arrange
	req := fetcher.Request{
		Method: "GET",
		Path:   "/user",
		Cookies: map[string]string{
			handlers.LoginTokenCookieName: loginToken,
		},
	}

	// Act
	resp, err := f.Fetch(req)

	// Assert
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.LessOrEqual(t, resp.StatusCode, 209)
	respBody := handlers.GetUserResBody{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	require.NoError(t, err)
	assert.Equal(t, userName, respBody.User.Name, "get user response should contain user name")
}

func testUpdateUser(t *testing.T, f *fetcher.Fetcher, loginToken string) {
	// Arrange
	req := fetcher.Request{
		Method: "PATCH",
		Path:   "/user",
		Body: handlers.UpdateUserReqBody{
			Data: handlers.UpdateUserReqBodyData{
				Name: userNameNew,
			},
		},
		Cookies: map[string]string{
			handlers.LoginTokenCookieName: loginToken,
		},
	}

	// Act
	resp, err := f.Fetch(req)

	// Assert
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.LessOrEqual(t, resp.StatusCode, 209)
}

func TestAuth(t *testing.T) {
	godotenv.Load("../.env")
	f := fetcher.New("http://localhost:1323")

	// Clear database
	clearDatabase(t)
	fmt.Println("✅ Cleared database")

	// Issue code
	codeId := testIssueCode(t, f)
	fmt.Println("✅ Issued code")
	fmt.Println("codeId:", codeId)

	// Verify code
	otpToken := testVerifyCode(t, f, codeId)
	fmt.Println("✅ Verified code")
	fmt.Println("otpToken:", otpToken)

	// Create user
	loginToken := testCreateUser(t, f, otpToken)
	fmt.Println("✅ Created user")
	fmt.Println("loginToken:", loginToken)

	// Get user
	testGetUser(t, f, loginToken)
	fmt.Println("✅ Got user")

	// Update user
	testUpdateUser(t, f, loginToken)
	fmt.Println("✅ Updated user")

	// Clear database
	clearDatabase(t)
	fmt.Println("✅ Cleared database")
}
