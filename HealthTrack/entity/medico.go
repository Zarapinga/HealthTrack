package entity

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Medico struct {
	Nome             string
	Email            string `gorm:"unique"`
	Idade            int
	DataDeNascimento string
	CPF              string `gorm:"unique"`
	Especialidade    string
	Salario          float64
	Senha            string
	ID               uuid.UUID `gorm:"primaryKey"`
}

func NewMedico(salario float64, nome, email, dataDeNascimento, cpf, especialidade, senha string, idade int) (*Medico, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	medico := &Medico{
		Nome:             nome,
		Email:            email,
		Idade:            idade,
		DataDeNascimento: dataDeNascimento,
		CPF:              cpf,
		Especialidade:    especialidade,
		Salario:          salario,
		Senha:            string(hash),
		ID:               uuid.New(),
	}
	err = medico.ValidateMedico()
	if err != nil {
		return nil, err
	}
	return medico, nil
}

func (m *Medico) ValidateMedico() error {
	if m.Nome == "" {
		return errors.New("Nome is required")
	}
	if m.Email == "" {
		return errors.New("Email is required")
	}
	if m.Idade <= 0 {
		return errors.New("Idade is invalid")
	}
	if m.DataDeNascimento == "" {
		return errors.New("DataDeNascimento is required")
	}
	if m.CPF == "" {
		return errors.New("CPF is required")
	}
	if m.Especialidade == "" {
		return errors.New("Especialidade is required")
	}
	if m.Salario <= 0 {
		return errors.New("Salario is invalid")
	}
	return nil
}
