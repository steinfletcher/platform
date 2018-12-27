package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	pw := "notTheRealPasswordLolz"

	hashed := HashPassword(pw)
	isValid := CheckPassword(pw, hashed)

	assert.True(t, isValid)
	assert.NotEqual(t, pw, hashed)
}
