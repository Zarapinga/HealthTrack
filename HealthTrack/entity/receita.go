package entity

import (
	"errors"
	"github.com/google/uuid"
)

type Receita struct {
	NomeDoRemedio   string
	MedicoEmail     string
	PacienteEmail   string
	AgendamentoDate string
	ID              uuid.UUID `gorm:"primaryKey"`
}

func NewReceita(nomeDoRemedio, medicoID, pacienteID, agendamentoID string) (*Receita, error) {
	receita := &Receita{
		NomeDoRemedio:   nomeDoRemedio,
		MedicoEmail:     medicoID,
		PacienteEmail:   pacienteID,
		AgendamentoDate: agendamentoID,
		ID:              uuid.New(),
	}
	err := receita.ValidateReceita()
	if err != nil {
		return nil, err
	}
	return receita, nil
}

func (r *Receita) ValidateReceita() error {
	if r.NomeDoRemedio == "" {
		return errors.New("NomeDoRemedio is required")
	}
	if r.MedicoEmail == "" {
		return errors.New("MedicoID is required")
	}
	if r.PacienteEmail == "" {
		return errors.New("PacienteID is required")
	}
	if r.AgendamentoDate == "" {
		return errors.New("AgendamentoID is required")
	}
	return nil
}
