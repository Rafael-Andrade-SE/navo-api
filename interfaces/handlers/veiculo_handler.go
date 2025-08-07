package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rafael-andrade-se/navo/api/domain/entities"
	"github.com/rafael-andrade-se/navo/api/domain/services"
	"github.com/rafael-andrade-se/navo/api/dto"
)

type VeiculoHandler struct {
	Service services.VeiculoService
}

func NewVeiculoHandler(service services.VeiculoService) VeiculoHandler {
	return VeiculoHandler{Service: service}
}

func RegisterVeiculoRoutes(rg *gin.RouterGroup, handler VeiculoHandler) {
	rg.POST("/veiculos/usuario/:idUsuario", handler.CriarVeiculo)
	rg.GET("/veiculos/usuario/:idUsuario", handler.ListarVeiculosPorUsuario)
	rg.GET("/veiculos/:id", handler.ObterVeiculoPorID)
	rg.PUT("/veiculos/:id", handler.AtualizarVeiculo)
	rg.DELETE("/veiculos/:id", handler.RemoverVeiculo)
}

func (h *VeiculoHandler) CriarVeiculo(c *gin.Context) {
	var input dto.CriarVeiculoDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos", "detalhes": err.Error()})
		return
	}

	idUsuario, err := strconv.Atoi(c.Param("idUsuario"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID de usuário inválido"})
		return
	}

	veiculo := entities.Veiculo{
		UsuarioID:     uint(idUsuario),
		Placa:         input.Placa,
		Marca:         input.Marca,
		Modelo:        input.Modelo,
		AnoFabricacao: input.AnoFabricacao,
		Cor:           input.Cor,
	}

	if err := h.Service.Criar(&veiculo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	resposta := dto.VeiculoRespostaDTO{
		ID:            veiculo.ID,
		Placa:         veiculo.Placa,
		Marca:         veiculo.Marca,
		Modelo:        veiculo.Modelo,
		AnoFabricacao: veiculo.AnoFabricacao,
		Cor:           veiculo.Cor,
	}

	c.JSON(http.StatusCreated, resposta)
}

func (h *VeiculoHandler) ListarVeiculosPorUsuario(c *gin.Context) {
	idUsuario, err := strconv.Atoi(c.Param("idUsuario"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID de usuário inválido"})
		return
	}

	veiculos, err := h.Service.ListarPorUsuario(uint(idUsuario))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	var resposta []dto.VeiculoRespostaDTO
	for _, v := range veiculos {
		resposta = append(resposta, dto.VeiculoRespostaDTO{
			ID:            v.ID,
			Placa:         v.Placa,
			Marca:         v.Marca,
			Modelo:        v.Modelo,
			AnoFabricacao: v.AnoFabricacao,
			Cor:           v.Cor,
		})
	}

	c.JSON(http.StatusOK, resposta)
}

func (h *VeiculoHandler) ObterVeiculoPorID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	veiculo, err := h.Service.BuscarPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Veículo não encontrado"})
		return
	}

	resposta := dto.VeiculoRespostaDTO{
		ID:            veiculo.ID,
		Placa:         veiculo.Placa,
		Marca:         veiculo.Marca,
		Modelo:        veiculo.Modelo,
		AnoFabricacao: veiculo.AnoFabricacao,
		Cor:           veiculo.Cor,
	}

	c.JSON(http.StatusOK, resposta)
}

func (h *VeiculoHandler) AtualizarVeiculo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	var input dto.AtualizarVeiculoDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido", "detalhes": err.Error()})
		return
	}

	veiculo := entities.Veiculo{
		ID:            uint(id),
		Placa:         input.Placa,
		Marca:         input.Marca,
		Modelo:        input.Modelo,
		AnoFabricacao: input.AnoFabricacao,
		Cor:           input.Cor,
	}

	if err := h.Service.Atualizar(&veiculo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	resposta := dto.VeiculoRespostaDTO{
		ID:            veiculo.ID,
		Placa:         veiculo.Placa,
		Marca:         veiculo.Marca,
		Modelo:        veiculo.Modelo,
		AnoFabricacao: veiculo.AnoFabricacao,
		Cor:           veiculo.Cor,
	}

	c.JSON(http.StatusOK, resposta)
}

func (h *VeiculoHandler) RemoverVeiculo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	if err := h.Service.Remover(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Veículo removido com sucesso"})
}
