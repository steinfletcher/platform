package user

import (
	"github.com/steinfletcher/platform/shared/crypto"
	"github.com/steinfletcher/platform/shared/errors"
)

type Session struct {
	Username string
}

type GetSession func(token string) (*Session, *errors.Error)

func NewGetSession(verifyJWT crypto.VerifyJWT) GetSession {
	return func(token string) (*Session, *errors.Error) {
		claims, err := verifyJWT(token)
		if err != nil {
			return nil, errors.Unauthorized("you are not authorized")
		}
		return &Session{Username: claims["sub"].(string)}, nil
	}
}
