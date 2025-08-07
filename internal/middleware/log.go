package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rafael-andrade-se/navo/api/pkg/logger"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		inicio := time.Now()

		// Processar a requisição
		c.Next()

		duracao := time.Since(inicio)

		logger.Log.Info("Requisição recebida",
			"método", c.Request.Method,
			"rota", c.FullPath(),
			"status", c.Writer.Status(),
			"tempo", duracao.String(),
			"ip", c.ClientIP(),
		)
	}
}
