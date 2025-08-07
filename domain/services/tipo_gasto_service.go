package services

import (
	"errors"

	"github.com/rafael-andrade-se/navo/api/domain/entities"
	"github.com/rafael-andrade-se/navo/api/domain/repositories"
)

type TipoGastoService interface {
	Criar(tipo *entities.TipoGasto) error
	ListarTodos() ([]entities.TipoGasto, error)
	BuscarPorID(id uint) (*entities.TipoGasto, error)
	Atualizar(tipo *entities.TipoGasto) error
	Remover(id uint) error
}

type tipoGastoService struct {
	repo repositories.TipoGastoRepository
}

func NovoTipoGastoService(r repositories.TipoGastoRepository) TipoGastoService {
	return &tipoGastoService{repo: r}
}

func (s *tipoGastoService) Criar(tipo *entities.TipoGasto) error {
	if tipo.Nome == "" {
		return errors.New("nome é obrigatório")
	}
	return s.repo.Salvar(tipo)
}

func (s *tipoGastoService) ListarTodos() ([]entities.TipoGasto, error) {
	return s.repo.BuscarTodos()
}

func (s *tipoGastoService) BuscarPorID(id uint) (*entities.TipoGasto, error) {
	return s.repo.BuscarPorID(id)
}

func (s *tipoGastoService) Atualizar(tipo *entities.TipoGasto) error {
	return s.repo.Atualizar(tipo)
}

func (s *tipoGastoService) Remover(id uint) error {
	return s.repo.Remover(id)
}
