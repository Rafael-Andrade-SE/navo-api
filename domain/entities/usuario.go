package entities

type Usuario struct {
	ID    uint    `gorm:"primaryKey"`
	Nome  string  `gorm:"size:100;not null"`
	Email string  `gorm:"size:150;not null;unique"`
	Senha string  `gorm:"not null"`
	Role  *string `gorm:"size:50;default:'motorista'"`
}

func (Usuario) TableName() string {
	return "usuarios"
}
