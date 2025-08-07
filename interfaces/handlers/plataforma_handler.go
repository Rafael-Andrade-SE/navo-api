package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rafael-andrade-se/navo/api/domain/entities"
	"github.com/rafael-andrade-se/navo/api/domain/services"
	"github.com/rafael-andrade-se/navo/api/dto"
)

type PlataformaHandler struct {
	Service services.PlataformaService
}

func NewPlataformaHandler(service services.PlataformaService) PlataformaHandler {
	return PlataformaHandler{Service: service}
}

func (h *PlataformaHandler) ListarPlataformas(c *gin.Context) {
	plataformas, err := h.Service.ListarTodas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	var resposta []dto.PlataformaRespostaDTO
	for _, p := range plataformas {
		resposta = append(resposta, dto.PlataformaRespostaDTO{
			ID:   p.ID,
			Nome: p.Nome,
		})
	}

	c.JSON(http.StatusOK, resposta)
}

func (h *PlataformaHandler) CriarPlataforma(c *gin.Context) {
	var input dto.CriarPlataformaDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos", "detalhes": err.Error()})
		return
	}

	plataforma := entities.Plataforma{
		Nome: input.Nome,
	}

	if err := h.Service.Criar(&plataforma); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	resposta := dto.PlataformaRespostaDTO{
		ID:   plataforma.ID,
		Nome: plataforma.Nome,
	}

	c.JSON(http.StatusCreated, resposta)
}

func (h *PlataformaHandler) ObterPlataformaPorID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	plataforma, err := h.Service.BuscarPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Plataforma não encontrada"})
		return
	}

	resposta := dto.PlataformaRespostaDTO{
		ID:   plataforma.ID,
		Nome: plataforma.Nome,
	}

	c.JSON(http.StatusOK, resposta)
}

func (h *PlataformaHandler) AtualizarPlataforma(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input dto.AtualizarPlataformaDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}

	plataforma := entities.Plataforma{
		ID:   uint(id),
		Nome: input.Nome,
	}

	if err := h.Service.Atualizar(&plataforma); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	resposta := dto.PlataformaRespostaDTO{
		ID:   plataforma.ID,
		Nome: plataforma.Nome,
	}

	c.JSON(http.StatusOK, resposta)
}

func (h *PlataformaHandler) RemoverPlataforma(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.Remover(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensagem": "Plataforma removida com sucesso"})
}
