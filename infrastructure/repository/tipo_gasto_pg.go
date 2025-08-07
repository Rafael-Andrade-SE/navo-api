package repository

import (
	"github.com/rafael-andrade-se/navo/api/domain/entities"
	"github.com/rafael-andrade-se/navo/api/domain/repositories"
	"gorm.io/gorm"
)

type tipoGastoRepositoryPG struct {
	db *gorm.DB
}

func NovoTipoGastoRepositoryPG(db *gorm.DB) repositories.TipoGastoRepository {
	return &tipoGastoRepositoryPG{db}
}

func (r *tipoGastoRepositoryPG) Salvar(tipo *entities.TipoGasto) error {
	return r.db.Create(tipo).Error
}

func (r *tipoGastoRepositoryPG) BuscarPorID(id uint) (*entities.TipoGasto, error) {
	var tipo entities.TipoGasto
	err := r.db.First(&tipo, id).Error
	if err != nil {
		return nil, err
	}
	return &tipo, nil
}

func (r *tipoGastoRepositoryPG) BuscarTodos() ([]entities.TipoGasto, error) {
	var tipos []entities.TipoGasto
	err := r.db.Find(&tipos).Error
	return tipos, err
}

func (r *tipoGastoRepositoryPG) Atualizar(tipo *entities.TipoGasto) error {
	return r.db.Save(tipo).Error
}

func (r *tipoGastoRepositoryPG) Remover(id uint) error {
	return r.db.Delete(&entities.TipoGasto{}, id).Error
}
