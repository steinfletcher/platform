package x

import (
	"github.com/gin-gonic/gin"
	"github.com/steinfletcher/platform/shared/crypto"
	"github.com/steinfletcher/platform/shared/errors"
	"github.com/steinfletcher/platform/shared/user"
	"net/http"
	"sync"
)

var contextPool sync.Pool

type Handler func(*Context)

type Context struct {
	*gin.Context

	session   *user.Session
	verifyJWT crypto.VerifyJWT
}

func (c *Context) Session() *user.Session {
	if c.session == nil {
		sessionToken, err := c.Cookie("Session-Token")
		if sessionToken == "" || err != nil {
			return nil
		}

		claims, err := c.verifyJWT(sessionToken)
		if err != nil {
			return nil
		}

		c.session = &user.Session{Username: claims["sub"].(string)}
		return c.session
	}

	return c.session
}

func acquireContext(ctx *gin.Context) *Context {
	v := contextPool.Get()
	var myContext *Context
	if v == nil {
		myContext = &Context{}
	} else {
		myContext = v.(*Context)
	}
	myContext.Reset(ctx)
	return myContext
}

func releaseContext(m *Context) {
	contextPool.Put(m)
}

func (h Handler) Serve(ctx *gin.Context) {
	myContext := acquireContext(ctx)
	h(myContext)
	releaseContext(myContext)
}

func Adapt(handler Handler, sessionSecret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		myContext := acquireContext(ctx)
		myContext.verifyJWT = crypto.NewVerifyJWT(sessionSecret)

		handler(myContext)

		releaseContext(myContext)
	}
}

func (c *Context) Reset(ctx *gin.Context) {
	c.Context = ctx
	c.session = nil
	c.verifyJWT = nil
}

func (c *Context) Err(err *errors.Error) {
	c.JSON(err.StatusCode, err)
}

func (c *Context) Unauthorized() {
	err := errors.Unauthorized("You must be logged in to perform this action")
	c.JSON(err.StatusCode, err)
}

func (c *Context) OK() {
	c.AbortWithStatus(http.StatusOK)
}
