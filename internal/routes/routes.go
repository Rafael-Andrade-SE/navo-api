package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafael-andrade-se/navo/api/interfaces/handlers"
)

func RegisterRoutes(router *gin.Engine, usuarioHandler handlers.UsuarioHandler, authHandler handlers.AuthHandler) {
	api := router.Group("/api/v1")

	// Rota de autenticação (login)
	api.POST("/login", authHandler.Login)

	// Rotas de usuários
	usuarios := api.Group("/usuarios")
	handlers.RegisterUsuarioRoutes(usuarios, usuarioHandler)
}
