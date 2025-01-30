package database

import (
	"github.com/Zarapinga/HealthTrack/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Paciente struct {
	DB *gorm.DB
}

func NewPaciente(db *gorm.DB) *Paciente {
	return &Paciente{DB: db}
}

func (p *Paciente) CreatePaciente(paciente *entity.Paciente) error {
	return p.DB.Create(paciente).Error
}

func (p *Paciente) FindAllPaciente() ([]entity.Paciente, error) {
	var pacientes []entity.Paciente
	err := p.DB.Find(&pacientes).Error
	return pacientes, err
}

func (p *Paciente) FindByPacienteEmail(email string) (*entity.Paciente, error) {
	var paciente entity.Paciente
	err := p.DB.First(&paciente, "email = ?", email).Error
	return &paciente, err
}

func (p *Paciente) FindByPacienteId(id uuid.UUID) (*entity.Paciente, error) {
	var paciente entity.Paciente
	err := p.DB.First(&paciente, "id = ?", id).Error
	return &paciente, err
}

func (p *Paciente) UpdatePaciente(paciente *entity.Paciente) error {
	_, err := p.FindByPacienteEmail(paciente.Email)
	if err != nil {
		return err
	}
	return p.DB.Save(paciente).Error
}

func (p *Paciente) DeletePaciente(email string) error {
	paciente, err := p.FindByPacienteEmail(email)
	if err != nil {
		return err
	}
	return p.DB.Delete(paciente).Error
}
