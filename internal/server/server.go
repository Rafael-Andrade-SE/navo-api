package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rafael-andrade-se/navo/api/internal/middleware"
)

func SetupMiddleware(router *gin.Engine) {
	// Logger estruturado com slog
	router.Use(middleware.LoggerMiddleware())

	// Recovery
	router.Use(gin.Recovery())

	// CORS
	router.Use(cors.Default())
}
