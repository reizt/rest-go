package id

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateCode() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}
