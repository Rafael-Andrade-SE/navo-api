package repositories

import "github.com/rafael-andrade-se/navo/api/domain/entities"

type UsuarioRepository interface {
	Salvar(usuario *entities.Usuario) error
	BuscarPorID(id uint) (*entities.Usuario, error)
	BuscarTodos() ([]entities.Usuario, error)
	Atualizar(usuario *entities.Usuario) error
	Remover(id uint) error
	BuscarPorEmail(email string) (*entities.Usuario, error) // para login ou verificação
}
