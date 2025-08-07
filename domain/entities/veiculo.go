package entities

import "time"

type Veiculo struct {
	ID            uint   `gorm:"primaryKey"`
	UsuarioID     uint   `gorm:"not null"`
	Placa         string `gorm:"size:10;not null;unique"`
	Marca         string `gorm:"size:50;not null"`
	Modelo        string `gorm:"size:50;not null"`
	AnoFabricacao int    `gorm:"not null"`
	Cor           string `gorm:"size:30;not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (Veiculo) TableName() string {
	return "veiculos"
}
