package e2e

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/reizt/rest-go/e2e/fetcher"
	"github.com/reizt/rest-go/handlers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	userName     = "John"
	userNameNew  = "Jane"
	userEmail    = "john@example.com"
	userPassword = "password"
)

func testIssueCode(t *testing.T, f *fetcher.Fetcher) string {
	// Arrange
	req := fetcher.Request{
		Method: "POST",
		Path:   "/auth/code/issue",
		Body: handlers.IssueCodeReqBody{
			Email:  userEmail,
			Action: "create-user",
		},
		Headers: map[string]string{},
	}

	// Act
	resp, err := f.Fetch(req)

	// Assert
	require.NoError(t, err)
	assert.LessOrEqual(t, resp.StatusCode, 209)
	issueCodeResBody := handlers.IssueCodeResBody{}
	err = json.NewDecoder(resp.Body).Decode(&issueCodeResBody)
	require.NoError(t, err)
	assert.NotEmpty(t, issueCodeResBody.CodeId, "issue code response should contain code id")

	return issueCodeResBody.CodeId
}

func testVerifyCode(t *testing.T, f *fetcher.Fetcher, codeId string) string {
	// Arrange
	req := fetcher.Request{
		Method: "POST",
		Path:   "/auth/code/verify",
		Body: handlers.VerifyCodeReqBody{
			CodeId: codeId,
		},
		Headers: map[string]string{},
	}

	// Act
	resp, err := f.Fetch(req)

	// Assert
	require.NoError(t, err)
	assert.LessOrEqual(t, resp.StatusCode, 209)
	assert.NotEmpty(t, resp.Headers[handlers.OTPTokenCookieName], "verify code response should contain token cookie")

	return resp.Headers[handlers.OTPTokenCookieName]
}

func testCreateUser(t *testing.T, f *fetcher.Fetcher, otpToken string) string {
	// Arrange
	req := fetcher.Request{
		Method: "POST",
		Path:   "/auth/create-user",
		Body: handlers.CreateUserReqBody{
			Token:    otpToken,
			Name:     userName,
			Password: userPassword,
		},
		Headers: map[string]string{},
	}

	// Act
	resp, err := f.Fetch(req)

	// Assert
	require.NoError(t, err)
	assert.LessOrEqual(t, resp.StatusCode, 209)
	assert.NotEmpty(t, resp.Cookies[handlers.LoginTokenCookieName], "create user response should contain login token cookie")

	return resp.Cookies[handlers.LoginTokenCookieName]
}

func testGetUser(t *testing.T, f *fetcher.Fetcher, loginToken string) {
	// Arrange
	req := fetcher.Request{
		Method: "GET",
		Path:   "/user",
		Headers: map[string]string{
			"Cookie": (&http.Cookie{
				Name:     handlers.LoginTokenCookieName,
				Value:    loginToken,
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteStrictMode,
			}).String(),
		},
	}

	// Act
	resp, err := f.Fetch(req)

	// Assert
	require.NoError(t, err)
	assert.LessOrEqual(t, resp.StatusCode, 209)
	getUserResBody := handlers.GetUserResBody{}
	err = json.NewDecoder(resp.Body).Decode(&getUserResBody)
	require.NoError(t, err)
	assert.Equal(t, userName, getUserResBody.User.Name, "get user response should contain user name")
}

func testUpdateUser(t *testing.T, f *fetcher.Fetcher, loginToken string) {
	// Arrange
	req := fetcher.Request{
		Method: "PATCH",
		Path:   "/user",
		Body: handlers.UpdateUserReqBody{
			Token: loginToken,
			Data: handlers.UpdateUserReqBodyData{
				Name: userNameNew,
			},
		},
		Headers: map[string]string{},
	}

	// Act
	resp, err := f.Fetch(req)

	// Assert
	require.NoError(t, err)
	assert.LessOrEqual(t, resp.StatusCode, 209)
}

func TestAuth(t *testing.T) {
	f := fetcher.New("http://localhost:1323")

	// Issue code
	codeId := testIssueCode(t, f)

	// Verify code
	otpToken := testVerifyCode(t, f, codeId)

	// Create user
	loginToken := testCreateUser(t, f, otpToken)

	// Get user
	testGetUser(t, f, loginToken)

	// Update user
	testUpdateUser(t, f, loginToken)
}
