package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafael-andrade-se/navo/api/interfaces/handlers"
)

func RegisterUsuarioRoutes(rg *gin.RouterGroup, handler handlers.UsuarioHandler) {
	rg.GET("/", handler.ListarUsuarios)
	rg.POST("/", handler.CriarUsuario)
	rg.GET("/:id", handler.ObterUsuarioPorID)
	rg.PUT("/:id", handler.AtualizarUsuario)
	rg.DELETE("/:id", handler.RemoverUsuario)
}
