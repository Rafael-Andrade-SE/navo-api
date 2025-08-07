package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafael-andrade-se/navo/api/interfaces/handlers"
)

func RegisterPlataformaRoutes(rg *gin.RouterGroup, handler handlers.PlataformaHandler) {
	rg.GET("/", handler.ListarPlataformas)
	rg.POST("/", handler.CriarPlataforma)
	rg.GET("/:id", handler.ObterPlataformaPorID)
	rg.PUT("/:id", handler.AtualizarPlataforma)
	rg.DELETE("/:id", handler.RemoverPlataforma)
}
