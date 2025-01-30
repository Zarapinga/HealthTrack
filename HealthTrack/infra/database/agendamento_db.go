package database

import (
	"github.com/Zarapinga/HealthTrack/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Agendamento struct {
	DB *gorm.DB
}

func NewAgendamento(db *gorm.DB) *Agendamento {
	return &Agendamento{DB: db}
}

func (a *Agendamento) CreateAgendamento(agendamento *entity.Agendamento) error {
	return a.DB.Create(agendamento).Error
}

func (a *Agendamento) FindAllAgendamentoByMedicoID(medicoID uuid.UUID) ([]entity.Agendamento, error) {
	var agendamentos []entity.Agendamento
	err := a.DB.Where("id = ?", medicoID).Find(&agendamentos).Error
	return agendamentos, err
}

func (a *Agendamento) FindAllAgendamentoByPacienteID(pacienteID uuid.UUID) ([]entity.Agendamento, error) {
	var agendamentos []entity.Agendamento
	err := a.DB.Where("id = ?", pacienteID).Find(&agendamentos).Error
	if err != nil {
		return nil, err
	}
	return agendamentos, nil
}

func (a *Agendamento) FindAllAgendamento() ([]entity.Agendamento, error) {
	var agendamento []entity.Agendamento
	err := a.DB.Find(&agendamento).Error
	if err != nil {
		return nil, err
	}
	return agendamento, nil
}

func (a *Agendamento) FindByAgendamentoID(agendamentoID uuid.UUID) (*entity.Agendamento, error) {
	var agendamento entity.Agendamento
	err := a.DB.First("id = ?", agendamentoID).Find(&agendamento).Error
	if err != nil {
		return nil, err
	}
	return &agendamento, nil
}

func (a *Agendamento) UpdateAgendamento(agendamento *entity.Agendamento) error {
	_, err := a.FindByAgendamentoID(agendamento.ID)
	if err != nil {
		return err
	}
	return a.DB.Save(agendamento).Error
}

func (a *Agendamento) DeleteAgendamento(agendamentoID uuid.UUID) error {
	agendamento, err := a.FindByAgendamentoID(agendamentoID)
	if err != nil {
		return err
	}
	return a.DB.Delete(agendamento).Error
}
