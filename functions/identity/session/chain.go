package session

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/steinfletcher/platform/functions/identity/config"
	"github.com/steinfletcher/platform/shared/crypto"
	"github.com/steinfletcher/platform/shared/x"
)

func Chain(conf *config.Config, db *sqlx.DB) gin.HandlerFunc {
	createJWT := crypto.NewCreateJWT(conf.SessionSecret)
	selectUserPassword := NewSelectUserPassword(db)
	service := NewCreateSession(selectUserPassword, createJWT)

	handler := NewHandler(service, conf)

	return x.Adapt(handler, conf.SessionSecret)
}
