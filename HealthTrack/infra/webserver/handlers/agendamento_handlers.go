package handlers

import (
	"encoding/json"
	"github.com/Zarapinga/HealthTrack/entity"
	"github.com/Zarapinga/HealthTrack/infra/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"net/http"
)

type AgendamentoHandlers struct {
	AgendamentoDB database.AgendamentoInterface
}

func NewAgendamentoHandler(db database.AgendamentoInterface) *AgendamentoHandlers {
	return &AgendamentoHandlers{
		AgendamentoDB: db,
	}
}

func (a *AgendamentoHandlers) CreateAgendamentoHandler(w http.ResponseWriter, r *http.Request) {
	var receptor entity.Agendamento
	err := json.NewDecoder(r.Body).Decode(&receptor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}
	agendamento, err := entity.NewAgendamento(receptor.DataDoAgendamento, receptor.Valor, receptor.MedicoID, receptor.PacienteID)
	if err != nil {
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}
	err = a.AgendamentoDB.CreateAgendamento(agendamento)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (a *AgendamentoHandlers) GetAgendamentoByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uuID, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := a.AgendamentoDB.FindByAgendamentoID(uuID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (a *AgendamentoHandlers) FindByAgendamentoId(id uuid.UUID) bool {
	_, err := a.AgendamentoDB.FindByAgendamentoID(id)
	if err != nil {
		return false
	}
	return true
}

func (a *AgendamentoHandlers) GetAgendamentoByMedicoID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uuID, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := a.AgendamentoDB.FindAllAgendamentoByMedicoID(uuID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (a *AgendamentoHandlers) GetAgendamentoByPacienteID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uuID, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := a.AgendamentoDB.FindAllAgendamentoByPacienteID(uuID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (a *AgendamentoHandlers) UpdateAgendamento(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var agendamento entity.Agendamento
	err := json.NewDecoder(r.Body).Decode(&agendamento)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	agendamento.ID, err = uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = a.AgendamentoDB.FindByAgendamentoID(agendamento.ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = a.AgendamentoDB.UpdateAgendamento(&agendamento)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *AgendamentoHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uuID, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = a.AgendamentoDB.FindByAgendamentoID(uuID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = a.AgendamentoDB.DeleteAgendamento(uuID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
