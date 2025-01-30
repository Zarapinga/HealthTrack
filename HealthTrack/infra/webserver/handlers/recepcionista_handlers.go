package handlers

import (
	"github.com/Zarapinga/HealthTrack/entity"
	"github.com/Zarapinga/HealthTrack/infra/database"
	"html/template"
	"net/http"
	"strconv"
)

type RecepcionistaHandlers struct {
	RecepcionistaDB database.RecepcionistaInterface
}

func NewRecepcionistaHandler(db database.RecepcionistaInterface) *RecepcionistaHandlers {
	return &RecepcionistaHandlers{
		RecepcionistaDB: db,
	}
}

func (re *RecepcionistaHandlers) CadastrarRecepcionista(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templ, err := template.ParseFiles("templates/recepcionistas/cadastrar_recepcionista.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = templ.ExecuteTemplate(w, "cadastrar_recepcionista", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
		turno := r.FormValue("turno")
		senha := r.FormValue("senha")
		recepcionista, err := entity.NewRecepcionista(nome, email, dataDeNascimento, cpf, turno, senha, idade, salario)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = re.RecepcionistaDB.CreateRecepcionista(recepcionista)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func (re *RecepcionistaHandlers) ListarRecepcionista(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templ, err := template.ParseFiles("templates/recepcionistas/listar_recepcionista.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		recepcionistas, err := re.RecepcionistaDB.FindAllRecepcionista()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Recepcionistas []entity.Recepcionista
		}{
			Recepcionistas: recepcionistas,
		}
		err = templ.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (re *RecepcionistaHandlers) AtualizarRecepcionista(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templ, err := template.ParseFiles("templates/recepcionistas/atualizar_recepcionista.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		recepcionistas, err := re.RecepcionistaDB.FindAllRecepcionista()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Recepcionistas []entity.Recepcionista
		}{
			Recepcionistas: recepcionistas,
		}
		err = templ.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		email := r.FormValue("email")
		if email == "" {
			http.Error(w, "Email do recepcionista não informado", http.StatusBadRequest)
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
		turno := r.FormValue("turno")
		recepcionista, err := re.RecepcionistaDB.FindByRecepcionistaEmail(email)
		recepcionista.Nome = nome
		recepcionista.DataDeNascimento = dataDeNascimento
		recepcionista.Idade = idade
		recepcionista.Salario = salario
		recepcionista.Turno = turno
		err = re.RecepcionistaDB.UpdateRecepcionista(recepcionista)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func (re *RecepcionistaHandlers) DeleteRecepcionista(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templ, err := template.ParseFiles("templates/recepcionistas/listar_recepcionista.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		recepcionistas, err := re.RecepcionistaDB.FindAllRecepcionista()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Recepcionistas []entity.Recepcionista
		}{
			Recepcionistas: recepcionistas,
		}
		err = templ.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		email := r.FormValue("email")
		if email == "" {
			http.Error(w, "Email do recepcionista não informado", http.StatusBadRequest)
			return
		}

		err := re.RecepcionistaDB.DeleteRecepcionista(email)
		if err != nil {
			http.Error(w, "Erro ao deletar recepcionista", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
