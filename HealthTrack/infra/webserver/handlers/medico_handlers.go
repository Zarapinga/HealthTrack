package handlers

import (
	"github.com/Zarapinga/HealthTrack/entity"
	"github.com/Zarapinga/HealthTrack/infra/database"
	"github.com/google/uuid"
	"html/template"
	"net/http"
	"strconv"
)

type MedicoHandlers struct {
	MedicoDB database.MedicoInterface
}

func NewMedicoHandler(db database.MedicoInterface) *MedicoHandlers {
	return &MedicoHandlers{
		MedicoDB: db,
	}
}

func (m *MedicoHandlers) CadastrarMedico(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templ, err := template.ParseFiles("templates/medicos/cadastrar_medico.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = templ.ExecuteTemplate(w, "cadastrar_medico", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		nome := r.FormValue("nome")
		email := r.FormValue("email")
		idade, err := strconv.Atoi(r.FormValue("idade"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		dataDeNascimento := r.FormValue("dataDeNascimento")
		cpf := r.FormValue("cpf")
		salario, err := strconv.ParseFloat(r.FormValue("salario"), 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		especialidade := r.FormValue("especialidade")
		senha := r.FormValue("senha")
		medico, err := entity.NewMedico(salario, nome, email, dataDeNascimento, cpf, especialidade, senha, idade)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = m.MedicoDB.CreateMedico(medico)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func (m *MedicoHandlers) ListarMedico(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templ, err := template.ParseFiles("templates/medicos/listar_medico.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		medicos, err := m.MedicoDB.FindAllMedico()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Medicos []entity.Medico
		}{
			Medicos: medicos,
		}
		err = templ.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (m *MedicoHandlers) FindByMedicoId(id uuid.UUID) bool {
	_, err := m.MedicoDB.FindByMedicoId(id)
	if err != nil {
		return false
	}
	return true
}

func (m *MedicoHandlers) AtualizarMedico(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templ, err := template.ParseFiles("templates/medicos/atualizar_medico.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		medicos, err := m.MedicoDB.FindAllMedico()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Medicos []entity.Medico
		}{
			Medicos: medicos,
		}
		err = templ.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		email := r.FormValue("email")
		if email == "" {
			http.Error(w, "Email do médico não informado", http.StatusBadRequest)
			return
		}
		nome := r.FormValue("nome")
		dataDeNascimento := r.FormValue("dataDeNascimento")
		idade, err := strconv.Atoi(r.FormValue("idade"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		salario, err := strconv.ParseFloat(r.FormValue("salario"), 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		especialidade := r.FormValue("especialidade")
		medico, err := m.MedicoDB.FindByMedicoEmail(email)
		medico.Nome = nome
		medico.DataDeNascimento = dataDeNascimento
		medico.Idade = idade
		medico.Salario = salario
		medico.Especialidade = especialidade
		err = m.MedicoDB.UpdateMedico(medico)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func (m *MedicoHandlers) DeleteMedico(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templ, err := template.ParseFiles("templates/medicos/deletar_medico.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		medicos, err := m.MedicoDB.FindAllMedico()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Medicos []entity.Medico
		}{
			Medicos: medicos,
		}
		err = templ.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		email := r.FormValue("email")
		if email == "" {
			http.Error(w, "Email do medico não informado", http.StatusBadRequest)
			return
		}

		err := m.MedicoDB.DeleteMedico(email)
		if err != nil {
			http.Error(w, "Erro ao deletar medico", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
