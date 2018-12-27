package session

import (
	goErr "errors"
	"github.com/steinfletcher/platform/shared/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

const password = "$2a$08$3uRzHSNH1YI8GORQdgUgyufwJtyx5wR.DqiO1/Z7rK3UK37GzZjye"

func TestCreateSession_Success(t *testing.T) {
	selectPassword := func(username string) (string, error) {
		return password, nil
	}
	createJWT := func(subject string) (string, error) {
		return "sessionToken", nil
	}

	jwtToken, err := NewCreateSession(selectPassword, createJWT)("user", "a")

	assert.Nil(t, err)
	assert.Equal(t, "sessionToken", jwtToken)
}

func TestCreateSession_WrongPassword(t *testing.T) {
	selectPassword := func(username string) (string, error) {
		return password, nil
	}
	createJWT := func(subject string) (string, error) {
		return "sessionToken", nil
	}

	jwtToken, err := NewCreateSession(selectPassword, createJWT)("user", "wrong")

	assert.Empty(t, jwtToken)
	assert.Equal(t, errors.BadRequest("Invalid credentials"), err)
}

func TestCreateSession_ServerErrorIfDatabaseError(t *testing.T) {
	selectPassword := func(username string) (string, error) {
		return "", goErr.New("error")
	}

	jwtToken, err := NewCreateSession(selectPassword, nil)("user", "a")

	assert.Empty(t, jwtToken)
	assert.Equal(t, errors.ServerError(), err)
}

func TestCreateSession_ServerErrorIfSigningFailure(t *testing.T) {
	selectPassword := func(username string) (string, error) {
		return password, nil
	}
	createJWT := func(subject string) (string, error) {
		return "", goErr.New("error")
	}

	jwtToken, err := NewCreateSession(selectPassword, createJWT)("user", "a")

	assert.Empty(t, jwtToken)
	assert.Equal(t, errors.ServerError(), err)
}
