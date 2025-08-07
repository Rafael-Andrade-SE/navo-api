package services

import (
	"errors"

	"github.com/rafael-andrade-se/navo/api/domain/entities"
	"github.com/rafael-andrade-se/navo/api/domain/repositories"
)

type PlataformaService interface {
	Criar(plataforma *entities.Plataforma) error
	ListarTodas() ([]entities.Plataforma, error)
	BuscarPorID(id uint) (*entities.Plataforma, error)
	Atualizar(plataforma *entities.Plataforma) error
	Remover(id uint) error
}

type plataformaService struct {
	repo repositories.PlataformaRepository
}

func NovaPlataformaService(r repositories.PlataformaRepository) PlataformaService {
	return &plataformaService{repo: r}
}

func (s *plataformaService) Criar(plataforma *entities.Plataforma) error {
	if plataforma.Nome == "" {
		return errors.New("o nome da plataforma é obrigatório")
	}

	return s.repo.Salvar(plataforma)
}

func (s *plataformaService) ListarTodas() ([]entities.Plataforma, error) {
	return s.repo.BuscarTodas()
}

func (s *plataformaService) BuscarPorID(id uint) (*entities.Plataforma, error) {
	return s.repo.BuscarPorID(id)
}

func (s *plataformaService) Atualizar(plataforma *entities.Plataforma) error {
	return s.repo.Atualizar(plataforma)
}

func (s *plataformaService) Remover(id uint) error {
	return s.repo.Remover(id)
}
