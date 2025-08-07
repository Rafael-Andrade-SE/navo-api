package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafael-andrade-se/navo/api/interfaces/handlers"
)

func RegisterRoutes(
	router *gin.Engine,
	usuarioHandler handlers.UsuarioHandler,
	authHandler handlers.AuthHandler,
	veiculoHandler handlers.VeiculoHandler,
) {
	api := router.Group("/api/v1")

	RegisterAuthRoutes(api, authHandler)
	RegisterUsuarioRoutes(api.Group("/usuarios"), usuarioHandler)
	RegisterVeiculoRoutes(api.Group("/veiculos"), veiculoHandler)
}
