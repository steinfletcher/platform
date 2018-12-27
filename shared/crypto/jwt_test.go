package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var secretKey = "1234567890"

func TestCreateJWT_GeneratesAValidToken(t *testing.T) {
	token, err := NewCreateJWT(secretKey)("userID")

	assert.Nil(t, err)
	assert.NotEmpty(t, token)
	claims, err := NewVerifyJWT(secretKey)(token)
	assert.Nil(t, err)
	assert.Equal(t, "userID", claims["sub"])
}

func TestCreateJWT_CannotBeTamperedWith(t *testing.T) {
	tokenWithModifiedSubject := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcGlrZXkiOiIxMjM0NSIsImV4cCI6MTU0MzQwMTU2OSwiaWF0IjoxNTQzNDAxMjY5LCJzdWIiOiI1NDMyMiJ9.3rnlNdSSJpFRJ1fTRiFBA1oTpqYtRSZ9r9VVSE-H0bE"

	claims, err := NewVerifyJWT(secretKey)(tokenWithModifiedSubject)

	assert.Len(t, claims, 0)
	assert.NotNil(t, err)
}
