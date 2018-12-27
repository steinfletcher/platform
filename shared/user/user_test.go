package user

import (
	goErr "errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/steinfletcher/platform/shared/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSession_Success(t *testing.T) {
	verifyJWT := func(token string) (jwt.MapClaims, error) {
		return map[string]interface{}{"sub": "email@me.com"}, nil
	}

	user, err := NewGetSession(verifyJWT)("sometoken")

	assert.Nil(t, err)
	assert.Equal(t, "email@me.com", user.Username)
}

func TestGetSession_BadRequestIfJWTInvalid(t *testing.T) {
	verifyJWT := func(token string) (jwt.MapClaims, error) {
		return map[string]interface{}{}, goErr.New("invalid JWT")
	}

	user, err := NewGetSession(verifyJWT)("sometoken")

	assert.Nil(t, user)
	assert.Equal(t, errors.BadRequest("Invalid session token"), err)
}
