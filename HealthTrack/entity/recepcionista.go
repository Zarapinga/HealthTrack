package entity

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Recepcionista struct {
	Nome             string
	Email            string `gorm:"unique"`
	Idade            int
	DataDeNascimento string
	CPF              string `gorm:"unique"`
	Salario          float64
	Turno            string
	Senha            string
	ID               uuid.UUID `gorm:"primaryKey"`
}

func NewRecepcionista(nome, email, dataDeNascimento, cpf, turno, senha string, idade int, salario float64) (*Recepcionista, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	recepcionista := &Recepcionista{
		Nome:             nome,
		Email:            email,
		Idade:            idade,
		DataDeNascimento: dataDeNascimento,
		CPF:              cpf,
		Salario:          salario,
		Turno:            turno,
		Senha:            string(hash),
		ID:               uuid.New(),
	}
	err = recepcionista.ValidateRecepcionista()
	if err != nil {
		return nil, err
	}
	return recepcionista, nil
}

func (r *Recepcionista) ValidateRecepcionista() error {
	if r.Nome == "" {
		return errors.New("Nome is required")
	}
	if r.Email == "" {
		return errors.New("Email is required")
	}
	if r.Idade <= 0 {
		return errors.New("Idade is invalid")
	}
	if r.DataDeNascimento == "" {
		return errors.New("DataDeNascimento is required")
	}
	if r.CPF == "" {
		return errors.New("CPF is required")
	}
	if r.Turno == "" {
		return errors.New("Turno is required")
	}
	if r.Salario <= 0 {
		return errors.New("Salario is invalid")
	}
	return nil
}
