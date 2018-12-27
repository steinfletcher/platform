package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/steinfletcher/platform/functions/identity/config"
	"github.com/steinfletcher/platform/functions/identity/session"
	"github.com/steinfletcher/platform/shared/x"
	"net/http"
)

type App struct {
	Router *gin.Engine
	Conf   *config.Config
	DB     *sqlx.DB
}

func New() *App {
	conf := config.New()
	db := sqlx.MustConnect("postgres", conf.DBAddr)

	router := gin.Default()
	router.Use(x.Cors(conf.UIDomain))

	a := &App{
		Router: router,
		Conf:   conf,
		DB:     db,
	}

	a.Router.GET("/v1/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	a.registerRoutes()
	return a
}

func (a *App) registerRoutes() {
	a.Router.POST("/v1/session", session.Chain(a.Conf, a.DB))

	// FIXME do this generically
	a.Router.OPTIONS("/v1/session", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
	})
}

func (a *App) Start() {
	x.Start(a.Router, a.Conf.Port)
}
