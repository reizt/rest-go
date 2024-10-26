package id

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateId(t *testing.T) {
	id := GenerateId()
	fmt.Println(id)
	assert.NotEmpty(t, id)
}
