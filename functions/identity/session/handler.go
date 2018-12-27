package session

import (
	"github.com/steinfletcher/platform/functions/identity/config"
	"github.com/steinfletcher/platform/shared/errors"
	"github.com/steinfletcher/platform/shared/x"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewHandler(createSession CreateSession, conf *config.Config) x.Handler {
	return func(c *x.Context) {
		var json LoginRequest
		if err := c.ShouldBindJSON(&json); err != nil {
			c.Err(errors.InvalidRequestBody)
			return
		}

		jwtToken, err := createSession(json.Username, json.Password)
		if err != nil {
			c.Err(err)
			return
		}

		http.SetCookie(c.Writer, &http.Cookie{
			Name:   "Session-Token",
			Value:  jwtToken,
			Domain: "." + conf.UIDomain,
		})
		c.OK()
	}
}
