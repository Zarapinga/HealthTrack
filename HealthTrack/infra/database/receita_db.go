package database

import (
	"github.com/Zarapinga/HealthTrack/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Receita struct {
	DB *gorm.DB
}

func NewReceita(db *gorm.DB) *Receita {
	return &Receita{DB: db}
}

func (r *Receita) CreateReceita(receita *entity.Receita) error {
	return r.DB.Create(receita).Error
}

func (r *Receita) FindAllReceitaByMedicoID(medicoID uuid.UUID) ([]entity.Receita, error) {
	var receitas []entity.Receita
	err := r.DB.Where("id = ?", medicoID).Find(&receitas).Error
	return receitas, err
}

func (r *Receita) FindAllReceitaByPacienteID(pacienteID uuid.UUID) ([]entity.Receita, error) {
	var receitas []entity.Receita
	err := r.DB.Where("id = ?", pacienteID).Find(&receitas).Error
	if err != nil {
		return nil, err
	}
	return receitas, nil
}

func (r *Receita) FindAllReceitaByAgendamentoID(agendamentoID uuid.UUID) ([]entity.Receita, error) {
	var receitas []entity.Receita
	err := r.DB.First(&receitas, "id = ?", agendamentoID).Error
	if err != nil {
		return nil, err
	}
	return receitas, nil
}

func (r *Receita) FindByReceitaID(receitaID uuid.UUID) (*entity.Receita, error) {
	var receita entity.Receita
	err := r.DB.First(&receita, "id = ?", receitaID).Error
	if err != nil {
		return nil, err
	}
	return &receita, nil
}

func (r *Receita) FindAllReceita() ([]entity.Receita, error) {
	var receita []entity.Receita
	err := r.DB.Find(&receita).Error
	if err != nil {
		return nil, err
	}
	return receita, nil
}

func (r *Receita) UpdateReceita(receita *entity.Receita) error {
	_, err := r.FindByReceitaID(receita.ID)
	if err != nil {
		return err
	}
	return r.DB.Save(receita).Error
}

func (r *Receita) DeleteReceita(receitaID uuid.UUID) error {
	agendamento, err := r.FindByReceitaID(receitaID)
	if err != nil {
		return err
	}
	return r.DB.Delete(agendamento).Error
}
