package get_profile

import (
	"github.com/gin-gonic/gin"
	"github.com/steinfletcher/platform/functions/user/config"
	"github.com/steinfletcher/platform/shared/x"
)

func Chain(conf *config.Config) gin.HandlerFunc {
	return x.Adapt(NewHandler(), conf.SessionSecret)
}
