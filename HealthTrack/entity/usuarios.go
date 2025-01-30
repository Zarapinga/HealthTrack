package entity

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Usuario struct {
	Email string `gorm:"primaryKey"`
	Senha string
	Tipo  string
}

func NewUsuario(email, senha, tipo string) (*Usuario, error) {
	usuario := &Usuario{
		Email: email,
		Senha: senha,
		Tipo:  tipo,
	}
	err := usuario.ValidateUsuario()
	if err != nil {
		return nil, err
	}
	return usuario, nil
}

func (u *Usuario) ValidateUsuario() error {
	if u.Email == "" {
		return errors.New("Email is required")
	}
	if u.Senha == "" {
		return errors.New("Senha is required")
	}
	if u.Tipo == "" {
		return errors.New("Tipo is required")
	}
	return nil
}

func (u *Usuario) ValidarSenha(senha string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Senha), []byte(senha))
	return err == nil
}
