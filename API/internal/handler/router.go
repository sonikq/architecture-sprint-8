package handler

import (
	"backend-api/internal/config"
	"backend-api/internal/server/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	keycloak "github.com/tbaehler/gin-keycloak/pkg/ginkeycloak"
	"net/http"
)

func NewRouter(cfg config.Config) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	headers := []string{"Content-Type", "Content-Length",
		"Accept-Encoding", "X-CSRF-Token",
		"Authorization", "accept", "origin",
		"Cache-Control", "X-Requested-With", "Token"}

	corsCfg := cors.DefaultConfig()
	corsCfg.AllowAllOrigins = true
	corsCfg.AllowCredentials = true
	corsCfg.AllowHeaders = append(corsCfg.AllowHeaders, headers...)
	router.Use(cors.New(corsCfg))

	keycloakCfg := keycloak.BuilderConfig{
		Service: cfg.Keycloak.Client,
		Url:     cfg.Keycloak.URI,
		Realm:   cfg.Keycloak.Realm,
	}
	authorized := router.Group("/reports")
	authorized.Use(middleware.HasAuth())
	authorized.Use(
		keycloak.NewAccessBuilder(keycloakCfg).
			RestrictButForRealm(cfg.Keycloak.AllowedRole).
			Build(),
	)
	authorized.GET("/", GetReports)

	return router
}
