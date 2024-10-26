package id

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateCode(t *testing.T) {
	code, err := GenerateCode()
	assert.NoError(t, err)
	fmt.Println(code)
	assert.NotEmpty(t, code)
}
