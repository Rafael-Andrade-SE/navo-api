package repository

import (
	"github.com/rafael-andrade-se/navo/api/domain/entities"
	"github.com/rafael-andrade-se/navo/api/domain/repositories"
	"gorm.io/gorm"
)

type usuarioRepositoryPG struct {
	db *gorm.DB
}

func NovoUsuarioRepositoryPG(db *gorm.DB) repositories.UsuarioRepository {
	return &usuarioRepositoryPG{db}
}

func (r *usuarioRepositoryPG) Salvar(usuario *entities.Usuario) error {
	return r.db.Create(usuario).Error
}

func (r *usuarioRepositoryPG) BuscarPorID(id uint) (*entities.Usuario, error) {
	var usuario entities.Usuario
	err := r.db.First(&usuario, id).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *usuarioRepositoryPG) BuscarTodos() ([]entities.Usuario, error) {
	var usuarios []entities.Usuario
	err := r.db.Find(&usuarios).Error
	return usuarios, err
}

func (r *usuarioRepositoryPG) Atualizar(usuario *entities.Usuario) error {
	return r.db.Save(usuario).Error
}

func (r *usuarioRepositoryPG) Remover(id uint) error {
	return r.db.Delete(&entities.Usuario{}, id).Error
}

func (r *usuarioRepositoryPG) BuscarPorEmail(email string) (*entities.Usuario, error) {
	var usuario entities.Usuario
	err := r.db.Where("email = ?", email).First(&usuario).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}
