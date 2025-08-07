package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rafael-andrade-se/navo/api/domain/entities"
	"github.com/rafael-andrade-se/navo/api/domain/services"
	"github.com/rafael-andrade-se/navo/api/dto"
	"golang.org/x/crypto/bcrypt"
)

type UsuarioHandler struct {
	Service services.UsuarioService
}

func NewUsuarioHandler(service services.UsuarioService) UsuarioHandler {
	return UsuarioHandler{Service: service}
}

// ---------------------------------------------

func (h *UsuarioHandler) ListarUsuarios(c *gin.Context) {
	usuarios, err := h.Service.ListarTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	var resposta []dto.UsuarioRespostaDTO
	for _, u := range usuarios {
		role := "motorista"
		if u.Role != nil {
			role = *u.Role
		}

		resposta = append(resposta, dto.UsuarioRespostaDTO{
			ID:    u.ID,
			Nome:  u.Nome,
			Email: u.Email,
			Role:  role,
		})
	}

	c.JSON(http.StatusOK, resposta)
}

func (h *UsuarioHandler) CriarUsuario(c *gin.Context) {
	var input dto.CriarUsuarioDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos", "detalhes": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Senha), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao gerar hash da senha"})
		return
	}

	usuario := entities.Usuario{
		Nome:  input.Nome,
		Email: input.Email,
		Senha: string(hash),
	}

	if input.Role != "" {
		usuario.Role = &input.Role
	}

	if err := h.Service.Criar(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	role := "motorista"
	if usuario.Role != nil {
		role = *usuario.Role
	}

	resposta := dto.UsuarioRespostaDTO{
		ID:    usuario.ID,
		Nome:  usuario.Nome,
		Email: usuario.Email,
		Role:  role,
	}

	c.JSON(http.StatusCreated, resposta)
}

func (h *UsuarioHandler) ObterUsuarioPorID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	usuario, err := h.Service.BuscarPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Usuário não encontrado"})
		return
	}

	role := "motorista"
	if usuario.Role != nil {
		role = *usuario.Role
	}

	resposta := dto.UsuarioRespostaDTO{
		ID:    usuario.ID,
		Nome:  usuario.Nome,
		Email: usuario.Email,
		Role:  role,
	}

	c.JSON(http.StatusOK, resposta)
}

func (h *UsuarioHandler) AtualizarUsuario(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input dto.AtualizarUsuarioDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Senha), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao gerar hash da senha"})
		return
	}

	usuario := entities.Usuario{
		ID:    uint(id),
		Nome:  input.Nome,
		Email: input.Email,
		Senha: string(hash),
	}

	if input.Role != "" {
		usuario.Role = &input.Role
	}

	if err := h.Service.Atualizar(&usuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	role := "motorista"
	if usuario.Role != nil {
		role = *usuario.Role
	}

	resposta := dto.UsuarioRespostaDTO{
		ID:    usuario.ID,
		Nome:  usuario.Nome,
		Email: usuario.Email,
		Role:  role,
	}

	c.JSON(http.StatusOK, resposta)
}

func (h *UsuarioHandler) RemoverUsuario(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.Remover(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensagem": "Usuário removido com sucesso"})
}
