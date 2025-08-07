package dto

type CriarPlataformaDTO struct {
	Nome string `json:"nome" binding:"required"`
}

type AtualizarPlataformaDTO struct {
	Nome string `json:"nome" binding:"required"`
}

type PlataformaRespostaDTO struct {
	ID   uint   `json:"id"`
	Nome string `json:"nome"`
}
