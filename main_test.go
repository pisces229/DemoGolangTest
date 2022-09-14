package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Default(test *testing.T) {
	assert.Equal(test, "test", "test")
}
