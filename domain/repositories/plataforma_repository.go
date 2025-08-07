package repositories

import "github.com/rafael-andrade-se/navo/api/domain/entities"

type PlataformaRepository interface {
	Salvar(plataforma *entities.Plataforma) error
	BuscarPorID(id uint) (*entities.Plataforma, error)
	BuscarTodas() ([]entities.Plataforma, error)
	Atualizar(plataforma *entities.Plataforma) error
	Remover(id uint) error
}
