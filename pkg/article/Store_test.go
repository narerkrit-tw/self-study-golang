package article

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadAll(t *testing.T) {
	result, err := LoadAll()

	assert.Nil(t, err)
	assert.Len(t, result, 2)
}

func TestLoadCategory_Success(t *testing.T) {
	result, err := LoadCategory("x")

	assert.Nil(t, err)
	assert.Len(t, result, 1)
}

func TestLoadCategory_Error(t *testing.T) {
	result, err := LoadCategory("error")

	assert.NotNil(t, err)
	assert.Len(t, result, 0)
}