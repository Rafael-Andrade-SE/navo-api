package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rafael-andrade-se/navo/api/domain/entities"
	"github.com/rafael-andrade-se/navo/api/domain/services"
	"github.com/rafael-andrade-se/navo/api/dto"
)

type TipoGastoHandler struct {
	Service services.TipoGastoService
}

func NewTipoGastoHandler(service services.TipoGastoService) TipoGastoHandler {
	return TipoGastoHandler{Service: service}
}

func (h *TipoGastoHandler) ListarTiposGasto(c *gin.Context) {
	tipos, err := h.Service.ListarTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	var resposta []dto.TipoGastoRespostaDTO
	for _, t := range tipos {
		resposta = append(resposta, dto.TipoGastoRespostaDTO{
			ID:   t.ID,
			Nome: t.Nome,
		})
	}

	c.JSON(http.StatusOK, resposta)
}

func (h *TipoGastoHandler) CriarTipoGasto(c *gin.Context) {
	var input dto.CriarTipoGastoDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos", "detalhes": err.Error()})
		return
	}

	tipo := entities.TipoGasto{
		Nome: input.Nome,
	}

	if err := h.Service.Criar(&tipo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	resposta := dto.TipoGastoRespostaDTO{
		ID:   tipo.ID,
		Nome: tipo.Nome,
	}

	c.JSON(http.StatusCreated, resposta)
}

func (h *TipoGastoHandler) ObterTipoGastoPorID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tipo, err := h.Service.BuscarPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Tipo de gasto não encontrado"})
		return
	}

	resposta := dto.TipoGastoRespostaDTO{
		ID:   tipo.ID,
		Nome: tipo.Nome,
	}

	c.JSON(http.StatusOK, resposta)
}

func (h *TipoGastoHandler) AtualizarTipoGasto(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input dto.AtualizarTipoGastoDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}

	tipo := entities.TipoGasto{
		ID:   uint(id),
		Nome: input.Nome,
	}

	if err := h.Service.Atualizar(&tipo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	resposta := dto.TipoGastoRespostaDTO{
		ID:   tipo.ID,
		Nome: tipo.Nome,
	}

	c.JSON(http.StatusOK, resposta)
}

func (h *TipoGastoHandler) RemoverTipoGasto(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.Remover(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensagem": "Tipo de gasto removido com sucesso"})
}
