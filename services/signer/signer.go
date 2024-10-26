package signer

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/reizt/rest-go/iservices/isigner"
)

type service struct {
	privateKey []byte
	publicKey  []byte
}

func New(privateKey, publicKey string) isigner.Service {
	return &service{
		privateKey: []byte(privateKey),
		publicKey:  []byte(publicKey),
	}
}

func (s *service) Sign(json string, expiresIn time.Duration) (string, error) {
	// Construct claims
	claims := jwt.MapClaims{
		"exp":  time.Now().Add(expiresIn).Unix(),
		"data": json,
	}

	// Sign token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *service) Verify(tokenString string) (string, error) {
	// Verify token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.publicKey, nil
	})
	if err != nil {
		return "", err
	}

	// Get payload
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	data, ok := claims["data"].(string)
	if !ok {
		return "", fmt.Errorf("invalid data in token")
	}

	return data, nil
}
