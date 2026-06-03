package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordNotEmpty(t *testing.T) {

	password := "123456"

	assert.NotEmpty(t, password)
}
