package id

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateCode(t *testing.T) {
	code, err := GenerateCode()
	require.NoError(t, err)
	fmt.Println(code)
	assert.NotEmpty(t, code)
}
