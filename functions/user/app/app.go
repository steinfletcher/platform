package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/steinfletcher/platform/functions/user/config"
	"github.com/steinfletcher/platform/functions/user/get-profile"
	"github.com/steinfletcher/platform/shared/x"
)

type App struct {
	Router *gin.Engine
	Conf   *config.Config
	DB     *sqlx.DB
}

func New() *App {
	conf := config.New()
	fmt.Println(conf.DBAddr)
	db := sqlx.MustConnect("postgres", conf.DBAddr)

	router := gin.Default()
	router.Use(x.Cors(conf.UIDomain))
	application := &App{
		Router: router,
		Conf:   conf,
		DB:     db,
	}
	application.registerRoutes()
	return application
}

func (a *App) registerRoutes() {
	a.Router.GET("/v1/user", get_profile.Chain(a.Conf))
	// FIXME do this generically
	a.Router.OPTIONS("/v1/user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	a.Router.GET("/v1/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
	})
}

func (a *App) Start() {
	x.Start(a.Router, a.Conf.Port)
}
