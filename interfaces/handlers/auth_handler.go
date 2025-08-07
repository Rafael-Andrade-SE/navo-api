package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafael-andrade-se/navo/api/domain/services"
	"github.com/rafael-andrade-se/navo/api/dto"
	"github.com/rafael-andrade-se/navo/api/pkg/jwt"
)

type AuthHandler struct {
	Service services.UsuarioService
}

func NewAuthHandler(service services.UsuarioService) AuthHandler {
	return AuthHandler{Service: service}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input dto.LoginDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	usuario, err := h.Service.Login(input.Email, input.Senha)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "Usuário ou senha inválidos"})
		return
	}

	role := "motorista"
	if usuario.Role != nil {
		role = *usuario.Role
	}

	token, err := jwt.GerarToken(usuario.ID, usuario.Email, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"usuario": gin.H{
			"id":    usuario.ID,
			"nome":  usuario.Nome,
			"email": usuario.Email,
			"role":  role,
		},
	})
}
