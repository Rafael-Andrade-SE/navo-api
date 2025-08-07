package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafael-andrade-se/navo/api/interfaces/handlers"
)

func RegisterTipoGastoRoutes(rg *gin.RouterGroup, handler handlers.TipoGastoHandler) {
	rg.GET("/", handler.ListarTiposGasto)
	rg.POST("/", handler.CriarTipoGasto)
	rg.GET("/:id", handler.ObterTipoGastoPorID)
	rg.PUT("/:id", handler.AtualizarTipoGasto)
	rg.DELETE("/:id", handler.RemoverTipoGasto)
}
