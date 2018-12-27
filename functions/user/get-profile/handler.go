package get_profile

import (
	"github.com/steinfletcher/platform/shared/x"
	"net/http"
)

type UserResponse struct {
	Username string `json:"username"`
}

func NewHandler() x.Handler {
	return func(c *x.Context) {
		session := c.Session()
		if session == nil {
			c.Unauthorized()
			return
		}

		c.JSON(http.StatusOK, UserResponse{Username: session.Username})
	}
}
