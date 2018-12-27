package x

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
)

func Cors(domainSuffix string) gin.HandlerFunc {
	defaultConfig := cors.DefaultConfig()
	return cors.New(cors.Config{
		AllowHeaders:     defaultConfig.AllowHeaders,
		AllowMethods:     defaultConfig.AllowMethods,
		AllowCredentials: true,
		AllowOriginFunc:  allowOriginFunc(domainSuffix),
	})
}

func allowOriginFunc(domainSuffix string) func(string) bool {
	return func(origin string) bool {
		return strings.HasSuffix(origin, domainSuffix)
	}
}
