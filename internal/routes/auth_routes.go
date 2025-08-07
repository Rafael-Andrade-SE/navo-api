package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafael-andrade-se/navo/api/interfaces/handlers"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, handler handlers.AuthHandler) {
	rg.POST("/login", handler.Login)
}
