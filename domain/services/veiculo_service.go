// path: api/domain/services/veiculo_service.go

package services

import (
	"github.com/rafael-andrade-se/navo/api/domain/entities"
	"github.com/rafael-andrade-se/navo/api/domain/repositories"
)

type VeiculoService interface {
	Criar(veiculo *entities.Veiculo) error
	ListarPorUsuario(usuarioID uint) ([]entities.Veiculo, error)
	BuscarPorID(id uint) (*entities.Veiculo, error)
	Atualizar(veiculo *entities.Veiculo) error
	Remover(id uint) error
}

type veiculoService struct {
	repo repositories.VeiculoRepository
}

func NovoVeiculoService(r repositories.VeiculoRepository) VeiculoService {
	return &veiculoService{repo: r}
}

func (s *veiculoService) Criar(veiculo *entities.Veiculo) error {
	return s.repo.Salvar(veiculo)
}

func (s *veiculoService) ListarPorUsuario(usuarioID uint) ([]entities.Veiculo, error) {
	return s.repo.ListarPorUsuario(usuarioID)
}

func (s *veiculoService) BuscarPorID(id uint) (*entities.Veiculo, error) {
	return s.repo.BuscarPorID(id)
}

func (s *veiculoService) Atualizar(veiculo *entities.Veiculo) error {
	return s.repo.Atualizar(veiculo)
}

func (s *veiculoService) Remover(id uint) error {
	return s.repo.Remover(id)
}
