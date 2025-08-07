package repositories

import "github.com/rafael-andrade-se/navo/api/domain/entities"

type VeiculoRepository interface {
	Salvar(veiculo *entities.Veiculo) error
	BuscarPorID(id uint) (*entities.Veiculo, error)
	ListarPorUsuario(usuarioID uint) ([]entities.Veiculo, error)
	Atualizar(veiculo *entities.Veiculo) error
	Remover(id uint) error
}
