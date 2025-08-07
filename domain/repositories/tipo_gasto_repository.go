package repositories

import "github.com/rafael-andrade-se/navo/api/domain/entities"

type TipoGastoRepository interface {
	Salvar(tipo *entities.TipoGasto) error
	BuscarPorID(id uint) (*entities.TipoGasto, error)
	BuscarTodos() ([]entities.TipoGasto, error)
	Atualizar(tipo *entities.TipoGasto) error
	Remover(id uint) error
}
