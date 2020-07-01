package basic_syntax

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {
	result := Hello()
	expect := "a"
	assert.Equal(t, expect, result)
}