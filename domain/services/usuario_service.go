package services

import (
	"errors"

	"github.com/rafael-andrade-se/navo/api/domain/entities"
	"github.com/rafael-andrade-se/navo/api/domain/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UsuarioService interface {
	Criar(usuario *entities.Usuario) error
	ListarTodos() ([]entities.Usuario, error)
	BuscarPorID(id uint) (*entities.Usuario, error)
	Atualizar(usuario *entities.Usuario) error
	Remover(id uint) error
	Login(email, senha string) (*entities.Usuario, error)
}

type usuarioService struct {
	repo repositories.UsuarioRepository
}

func NovoUsuarioService(r repositories.UsuarioRepository) UsuarioService {
	return &usuarioService{repo: r}
}

func (s *usuarioService) Criar(usuario *entities.Usuario) error {
	// Regras de validação básica (exemplo)
	if usuario.Nome == "" || usuario.Email == "" || usuario.Senha == "" {
		return errors.New("nome, email e senha são obrigatórios")
	}

	return s.repo.Salvar(usuario)
}

func (s *usuarioService) ListarTodos() ([]entities.Usuario, error) {
	return s.repo.BuscarTodos()
}

func (s *usuarioService) BuscarPorID(id uint) (*entities.Usuario, error) {
	return s.repo.BuscarPorID(id)
}

func (s *usuarioService) Atualizar(usuario *entities.Usuario) error {
	return s.repo.Atualizar(usuario)
}

func (s *usuarioService) Remover(id uint) error {
	return s.repo.Remover(id)
}

func (s *usuarioService) Login(email, senha string) (*entities.Usuario, error) {
	usuario, err := s.repo.BuscarPorEmail(email)
	if err != nil {
		return nil, errors.New("usuário ou senha inválidos")
	}

	err = bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(senha))
	if err != nil {
		return nil, errors.New("usuário ou senha inválidos")
	}

	return usuario, nil
}
