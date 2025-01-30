package entity

import (
	"github.com/Zarapinga/HealthTrack/configs"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Superusuario struct {
	Email string
	Senha string
	ID    uuid.UUID `gorm:"primary_key"`
}

func NewSuperusuario() *Superusuario {
	//NewUsuario(configs.EmailSuperUsuario, CriptografarSenha(), "superusuario")
	return &Superusuario{
		Email: configs.EmailSuperUsuario,
		Senha: CriptografarSenha(),
	}
}

func CriptografarSenha() string {
	hash, err := bcrypt.GenerateFromPassword([]byte(configs.SenhaSuperUsuario), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
