package session

import (
	"github.com/apex/log"
	"github.com/steinfletcher/platform/shared/crypto"
	"github.com/steinfletcher/platform/shared/errors"
)

type JwtToken = string

type CreateSession func(username, password string) (JwtToken, *errors.Error)

func NewCreateSession(selectPassword SelectPassword, createJWT crypto.CreateJWT) CreateSession {
	return func(username, password string) (JwtToken, *errors.Error) {
		userPassword, err := selectPassword(username)
		if err != nil {
			log.Errorf("select password query failed. err: %s", err)
			return "", errors.ServerError()
		}

		if userPassword == "" || !crypto.CheckPassword(password, userPassword) {
			return "", errors.BadRequest("Invalid credentials")
		}

		token, err := createJWT(username)
		if err != nil {
			log.Errorf("failed to sign jwt. err: %s", err)
			return "", errors.ServerError()
		}
		return token, nil
	}
}
