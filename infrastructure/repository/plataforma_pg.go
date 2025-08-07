package repository

import (
	"github.com/rafael-andrade-se/navo/api/domain/entities"
	"github.com/rafael-andrade-se/navo/api/domain/repositories"
	"gorm.io/gorm"
)

type plataformaRepositoryPG struct {
	db *gorm.DB
}

func NovaPlataformaRepositoryPG(db *gorm.DB) repositories.PlataformaRepository {
	return &plataformaRepositoryPG{db}
}

func (r *plataformaRepositoryPG) Salvar(plataforma *entities.Plataforma) error {
	return r.db.Create(plataforma).Error
}

func (r *plataformaRepositoryPG) BuscarPorID(id uint) (*entities.Plataforma, error) {
	var plataforma entities.Plataforma
	err := r.db.First(&plataforma, id).Error
	if err != nil {
		return nil, err
	}
	return &plataforma, nil
}

func (r *plataformaRepositoryPG) BuscarTodas() ([]entities.Plataforma, error) {
	var plataformas []entities.Plataforma
	err := r.db.Find(&plataformas).Error
	return plataformas, err
}

func (r *plataformaRepositoryPG) Atualizar(plataforma *entities.Plataforma) error {
	return r.db.Save(plataforma).Error
}

func (r *plataformaRepositoryPG) Remover(id uint) error {
	return r.db.Delete(&entities.Plataforma{}, id).Error
}
