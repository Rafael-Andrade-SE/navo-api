package dto

type CriarVeiculoDTO struct {
	Placa         string `json:"placa" binding:"required"`
	Marca         string `json:"marca" binding:"required"`
	Modelo        string `json:"modelo" binding:"required"`
	AnoFabricacao int    `json:"ano_fabricacao" binding:"required"`
	Cor           string `json:"cor" binding:"required"`
}

type AtualizarVeiculoDTO struct {
	Placa         string `json:"placa" binding:"required"`
	Marca         string `json:"marca" binding:"required"`
	Modelo        string `json:"modelo" binding:"required"`
	AnoFabricacao int    `json:"ano_fabricacao" binding:"required"`
	Cor           string `json:"cor" binding:"required"`
}

type VeiculoRespostaDTO struct {
	ID            uint   `json:"id"`
	Placa         string `json:"placa"`
	Marca         string `json:"marca"`
	Modelo        string `json:"modelo"`
	AnoFabricacao int    `json:"ano_fabricacao"`
	Cor           string `json:"cor"`
}
