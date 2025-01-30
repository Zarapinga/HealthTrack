package handlers

import (
	"fmt"
	"github.com/Zarapinga/HealthTrack/entity"
	database "github.com/Zarapinga/HealthTrack/infra/database"
	"github.com/google/uuid"
	"html/template"
	"net/http"
	"strconv"
)

type PacienteHandler struct {
	PacienteDB database.PacienteInterface
}

func NewPacienteHandler(db database.PacienteInterface) *PacienteHandler {
	return &PacienteHandler{
		PacienteDB: db,
	}
}

func (p *PacienteHandler) CadastrarPaciente(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templ, err := template.ParseFiles("templates/pacientes/cadastrar_paciente.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = templ.ExecuteTemplate(w, "cadastrar_paciente", nil)
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
		senha := r.FormValue("senha")
		paciente, err := entity.NewPaciente(nome, email, dataDeNascimento, cpf, senha, idade)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = p.PacienteDB.CreatePaciente(paciente)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func (p *PacienteHandler) ListarPaciente(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templ, err := template.ParseFiles("templates/pacientes/listar_paciente.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pacientes, err := p.PacienteDB.FindAllPaciente()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Pacientes []entity.Paciente
		}{
			Pacientes: pacientes,
		}
		err = templ.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (p *PacienteHandler) FindByPacienteId(id uuid.UUID) bool {
	_, err := p.PacienteDB.FindByPacienteId(id)
	if err != nil {
		return false
	}
	return true
}

func (p *PacienteHandler) AtualizarPaciente(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templ, err := template.ParseFiles("templates/pacientes/atualizar_paciente.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pacientes, err := p.PacienteDB.FindAllPaciente()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Pacientes []entity.Paciente
		}{
			Pacientes: pacientes,
		}
		err = templ.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		email := r.FormValue("email")
		nome := r.FormValue("nome")
		dataDeNascimento := r.FormValue("dataDeNascimento")
		idade, err := strconv.Atoi(r.FormValue("idade"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if email == "" {
			http.Error(w, "Email do paciente não informado", http.StatusBadRequest)
			return
		}
		paciente, err := p.PacienteDB.FindByPacienteEmail(email)
		if err != nil {
			http.Error(w, fmt.Sprintf("Paciente com email %s não encontrado", email), http.StatusNotFound)
			return
		}
		paciente.Nome = nome
		paciente.DataDeNascimento = dataDeNascimento
		paciente.Idade = idade
		err = p.PacienteDB.UpdatePaciente(paciente)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (p *PacienteHandler) DeletePacientes(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templ, err := template.ParseFiles("templates/pacientes/deletar_paciente.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pacientes, err := p.PacienteDB.FindAllPaciente()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Pacientes []entity.Paciente
		}{
			Pacientes: pacientes,
		}
		err = templ.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		email := r.FormValue("email")
		if email == "" {
			http.Error(w, "Email do paciente não informado", http.StatusBadRequest)
			return
		}
		_, err := p.PacienteDB.FindByPacienteEmail(email)
		if err != nil {
			http.Error(w, fmt.Sprintf("Paciente com email %s não encontrado", email), http.StatusNotFound)
			return
		}
		err = p.PacienteDB.DeletePaciente(email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
