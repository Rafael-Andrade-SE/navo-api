package entities

type TipoGasto struct {
	ID   uint   `gorm:"primaryKey"`
	Nome string `gorm:"size:100;not null;unique"`
}

func (TipoGasto) TableName() string {
	return "tipo_gastos"
}
