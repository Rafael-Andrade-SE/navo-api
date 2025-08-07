package dto

type CriarTipoGastoDTO struct {
	Nome string `json:"nome" binding:"required"`
}

type AtualizarTipoGastoDTO struct {
	Nome string `json:"nome" binding:"required"`
}

type TipoGastoRespostaDTO struct {
	ID   uint   `json:"id"`
	Nome string `json:"nome"`
}
