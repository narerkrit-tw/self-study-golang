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