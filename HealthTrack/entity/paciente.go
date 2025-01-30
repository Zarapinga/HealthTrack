package entity

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Paciente struct {
	Nome             string
	Email            string `gorm:"unique"`
	Idade            int
	DataDeNascimento string
	CPF              string `gorm:"unique"`
	Senha            string
	ID               uuid.UUID `gorm:"primaryKey"`
}

func NewPaciente(nome, email, dataDeNascimento, cpf, senha string, idade int) (*Paciente, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	paciente := &Paciente{
		Nome:             nome,
		Email:            email,
		Idade:            idade,
		DataDeNascimento: dataDeNascimento,
		CPF:              cpf,
		Senha:            string(hash),
		ID:               uuid.New(),
	}
	err = paciente.ValidatePaciente()
	if err != nil {
		return nil, err
	}
	return paciente, nil
}

func (p *Paciente) ValidatePaciente() error {
	if p.Nome == "" {
		return errors.New("Nome is required")
	}
	if p.Email == "" {
		return errors.New("Email is required")
	}
	if p.Idade <= 0 {
		return errors.New("Idade is invalid")
	}
	if p.DataDeNascimento == "" {
		return errors.New("DataDeNascimento is required")
	}
	if p.CPF == "" {
		return errors.New("CPF is required")
	}
	return nil
}
