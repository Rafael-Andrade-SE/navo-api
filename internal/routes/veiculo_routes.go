package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafael-andrade-se/navo/api/interfaces/handlers"
)

func RegisterVeiculoRoutes(rg *gin.RouterGroup, handler handlers.VeiculoHandler) {
	rg.POST("/usuario/:idUsuario", handler.CriarVeiculo)            // POST /api/v1/veiculos/usuario/:idUsuario
	rg.GET("/usuario/:idUsuario", handler.ListarVeiculosPorUsuario) // GET  /api/v1/veiculos/usuario/:idUsuario
	rg.GET("/:id", handler.ObterVeiculoPorID)                       // GET  /api/v1/veiculos/:id
	rg.PUT("/:id", handler.AtualizarVeiculo)                        // PUT  /api/v1/veiculos/:id
	rg.DELETE("/:id", handler.RemoverVeiculo)                       // DELETE /api/v1/veiculos/:id
}
