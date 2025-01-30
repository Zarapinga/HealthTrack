package database

import (
	"github.com/Zarapinga/HealthTrack/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Medico struct {
	DB *gorm.DB
}

func NewMedico(db *gorm.DB) *Medico {
	return &Medico{DB: db}
}

func (m *Medico) CreateMedico(medico *entity.Medico) error {
	return m.DB.Create(medico).Error
}

func (m *Medico) FindAllMedico() ([]entity.Medico, error) {
	var medico []entity.Medico
	err := m.DB.Find(&medico).Error
	return medico, err
}

func (m *Medico) FindByMedicoEmail(email string) (*entity.Medico, error) {
	var medico entity.Medico
	err := m.DB.First(&medico, "email = ?", email).Error
	return &medico, err
}

func (m *Medico) FindByMedicoId(id uuid.UUID) (*entity.Medico, error) {
	var medico entity.Medico
	err := m.DB.First(&medico, "id = ?", id).Error
	return &medico, err
}

func (m *Medico) UpdateMedico(medico *entity.Medico) error {
	_, err := m.FindByMedicoEmail(medico.Email)
	if err != nil {
		return err
	}
	return m.DB.Save(medico).Error
}

func (m *Medico) DeleteMedico(email string) error {
	medico, err := m.FindByMedicoEmail(email)
	if err != nil {
		return err
	}
	return m.DB.Delete(medico).Error
}
