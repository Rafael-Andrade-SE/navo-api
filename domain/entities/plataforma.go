package entities

type Plataforma struct {
	ID   uint   `gorm:"primaryKey"`
	Nome string `gorm:"size:100;not null;unique"`
}

func (Plataforma) TableName() string {
	return "plataformas"
}
