package database

import (
	"github.com/Zarapinga/HealthTrack/entity"
	"github.com/google/uuid"
)

type UsuarioInterface interface {
	CreateUsuario(user *entity.Usuario) error
	UpdateUsuario(user *entity.Usuario) error
	DeleteUsuario(email string) error
	FindByUsuarioEmail(email string) (*entity.Usuario, error)
}

type PacienteInterface interface {
	CreatePaciente(paciente *entity.Paciente) error
	FindAllPaciente() ([]entity.Paciente, error)
	FindByPacienteEmail(email string) (*entity.Paciente, error)
	FindByPacienteId(uuid uuid.UUID) (*entity.Paciente, error)
	UpdatePaciente(paciente *entity.Paciente) error
	DeletePaciente(email string) error
}

type MedicoInterface interface {
	CreateMedico(medico *entity.Medico) error
	FindAllMedico() ([]entity.Medico, error)
	FindByMedicoEmail(email string) (*entity.Medico, error)
	FindByMedicoId(id uuid.UUID) (*entity.Medico, error)
	UpdateMedico(medico *entity.Medico) error
	DeleteMedico(email string) error
}

type RecepcionistaInterface interface {
	CreateRecepcionista(recepcionista *entity.Recepcionista) error
	FindAllRecepcionista() ([]entity.Recepcionista, error)
	FindByRecepcionistaEmail(email string) (*entity.Recepcionista, error)
	UpdateRecepcionista(recepcionista *entity.Recepcionista) error
	DeleteRecepcionista(email string) error
}

type AgendamentoInterface interface {
	CreateAgendamento(agendamento *entity.Agendamento) error
	FindAllAgendamentoByMedicoID(medicoID uuid.UUID) ([]entity.Agendamento, error)
	FindAllAgendamentoByPacienteID(pacienteID uuid.UUID) ([]entity.Agendamento, error)
	FindByAgendamentoID(agendamentoID uuid.UUID) (*entity.Agendamento, error)
	UpdateAgendamento(agendamento *entity.Agendamento) error
	DeleteAgendamento(agendamentoID uuid.UUID) error
}

type ReceitaInterface interface {
	CreateReceita(receita *entity.Receita) error
	FindAllReceitaByMedicoID(medicoID uuid.UUID) ([]entity.Receita, error)
	FindAllReceitaByPacienteID(pacienteID uuid.UUID) ([]entity.Receita, error)
	FindAllReceitaByAgendamentoID(agendamentoID uuid.UUID) ([]entity.Receita, error)
	FindByReceitaID(receitaID uuid.UUID) (*entity.Receita, error)
	FindAllReceita() ([]entity.Receita, error)
	UpdateReceita(receita *entity.Receita) error
	DeleteReceita(receitaID uuid.UUID) error
}
