package signer

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/reizt/rest-go/iservices/isigner"
)

type service struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

func New() (isigner.Service, error) {
	privateKeyStr := os.Getenv("JWT_PRIVATE_KEY")
	publicKeyStr := os.Getenv("JWT_PUBLIC_KEY")
	if privateKeyStr == "" || publicKeyStr == "" {
		return nil, fmt.Errorf("JWT_PRIVATE_KEY and JWT_PUBLIC_KEY must be set")
	}

	privateKeyInfo, _ := pem.Decode([]byte(privateKeyStr))
	publicKeyInfo, _ := pem.Decode([]byte(publicKeyStr))

	privateKey, err := x509.ParsePKCS8PrivateKey(privateKeyInfo.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey, err := x509.ParsePKIXPublicKey(publicKeyInfo.Bytes)
	if err != nil {
		return nil, err
	}

	return &service{
		privateKey: privateKey.(*ecdsa.PrivateKey),
		publicKey:  publicKey.(*ecdsa.PublicKey),
	}, nil
}

func (s *service) Sign(json string, expiresIn time.Duration) (string, error) {
	// Construct claims
	claims := jwt.MapClaims{
		"exp":  time.Now().Add(expiresIn).Unix(),
		"data": json,
	}

	// Sign token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(s.privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *service) Verify(tokenString string) (string, error) {
	// Verify token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
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
