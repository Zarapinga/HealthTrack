package database

import (
	"github.com/Zarapinga/HealthTrack/entity"
	"gorm.io/gorm"
)

type Usuario struct {
	DB *gorm.DB
}

func (u *Usuario) CreateUsuario(usuario *entity.Usuario) error {
	return u.DB.Create(usuario).Error
}

func (u *Usuario) FindByUsuarioEmail(email string) (*entity.Usuario, error) {
	var usuario entity.Usuario
	err := u.DB.First(&usuario, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (u *Usuario) UpdateUsuario(usuario *entity.Usuario) error {
	_, err := u.FindByUsuarioEmail(usuario.Email)
	if err != nil {
		return err
	}
	return u.DB.Save(usuario).Error
}

func (u *Usuario) DeleteUsuario(email string) error {
	usuario, err := u.FindByUsuarioEmail(email)
	if err != nil {
		return err
	}
	return u.DB.Delete(usuario).Error
}
