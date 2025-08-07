package dto

type CriarUsuarioDTO struct {
	Nome  string `json:"nome" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required"`
	Role  string `json:"role" binding:"required"`
}

type AtualizarUsuarioDTO struct {
	Nome  string `json:"nome" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required"`
	Role  string `json:"role" binding:"required"`
}

type UsuarioRespostaDTO struct {
	ID    uint   `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
