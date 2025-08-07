package repository

import (
	"github.com/rafael-andrade-se/navo/api/domain/entities"
	"github.com/rafael-andrade-se/navo/api/domain/repositories"
	"gorm.io/gorm"
)

type veiculoRepositoryPG struct {
	db *gorm.DB
}

func NovoVeiculoRepositoryPG(db *gorm.DB) repositories.VeiculoRepository {
	return &veiculoRepositoryPG{db}
}

func (r *veiculoRepositoryPG) Salvar(veiculo *entities.Veiculo) error {
	return r.db.Create(veiculo).Error
}

func (r *veiculoRepositoryPG) BuscarPorID(id uint) (*entities.Veiculo, error) {
	var v entities.Veiculo
	err := r.db.First(&v, id).Error
	return &v, err
}

func (r *veiculoRepositoryPG) ListarPorUsuario(usuarioID uint) ([]entities.Veiculo, error) {
	var veiculos []entities.Veiculo
	err := r.db.Where("usuario_id = ?", usuarioID).Find(&veiculos).Error
	return veiculos, err
}

func (r *veiculoRepositoryPG) Atualizar(veiculo *entities.Veiculo) error {
	return r.db.Save(veiculo).Error
}

func (r *veiculoRepositoryPG) Remover(id uint) error {
	return r.db.Delete(&entities.Veiculo{}, id).Error
}
