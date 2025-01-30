package entity

import (
	"errors"
	"github.com/google/uuid"
)

type Agendamento struct {
	DataDoAgendamento string
	MedicoID          uuid.UUID
	PacienteID        uuid.UUID
	Valor             float64
	ID                uuid.UUID `gorm:"primaryKey"`
}

func NewAgendamento(data string, valor float64, medicoID, pacienteID uuid.UUID) (*Agendamento, error) {
	agendamento := &Agendamento{
		DataDoAgendamento: data,
		MedicoID:          medicoID,
		PacienteID:        pacienteID,
		Valor:             valor,
		ID:                uuid.New(),
	}
	err := agendamento.ValidateAgendamento()
	if err != nil {
		return nil, err
	}
	return agendamento, nil
}

func (a *Agendamento) ValidateAgendamento() error {
	if a.DataDoAgendamento == "" {
		return errors.New("DataDoAgendamento is required")
	}
	if a.MedicoID.String() == "" {
		return errors.New("MedicoID is required")
	}
	if a.PacienteID.String() == "" {
		return errors.New("PacienteID is required")
	}
	if a.Valor <= 0 {
		return errors.New("Valor is invalid")
	}
	return nil
}
