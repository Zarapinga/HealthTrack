package handlers

import (
	"github.com/Zarapinga/HealthTrack/entity"
	"github.com/Zarapinga/HealthTrack/infra/database"
	"github.com/google/uuid"
	"html/template"
	"net/http"
)

type ReceitaHandlers struct {
	ReceitaDB database.ReceitaInterface
}

func NewReceitaHandler(db database.ReceitaInterface) *ReceitaHandlers {
	return &ReceitaHandlers{
		ReceitaDB: db,
	}
}

func (re *ReceitaHandlers) CadastrarReceita(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templ, err := template.ParseFiles("templates/receitas/cadastrar_receita.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = templ.ExecuteTemplate(w, "cadastrar_receita", nil)
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
		nomeDoRemedio := r.FormValue("nomeDoRemedio")
		medicoId := r.FormValue("medicoId")
		mId, err := uuid.Parse(medicoId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if re.verficarMedico(mId) {
			pacienteId := r.FormValue("pacienteID")
			pId, err := uuid.Parse(medicoId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if re.verficarPaciente(pId) {
				agendamentoId := r.FormValue("agendamentoId")
				aId, err := uuid.Parse(agendamentoId)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				if re.verficarAgendamento(aId) {
					receita, err := entity.NewReceita(nomeDoRemedio, medicoId, pacienteId, agendamentoId)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					err = re.ReceitaDB.CreateReceita(receita)
					if err != nil {
						w.WriteHeader(http.StatusBadRequest)
						return
					}
				}
			}

			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func (re *ReceitaHandlers) ListarReceitas(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templ, err := template.ParseFiles("templates/receita/listar_receita.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		receita, err := re.ReceitaDB.FindAllReceita()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Receitas []entity.Receita
		}{
			Receitas: receita,
		}
		err = templ.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (re *ReceitaHandlers) AtualizarReceita(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templ, err := template.ParseFiles("templates/receita/atualizar_receita.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		receitas, err := re.ReceitaDB.FindAllReceita()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Receitas []entity.Receita
		}{
			Receitas: receitas,
		}
		err = templ.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		id := r.FormValue("id")
		if id == "" {
			http.Error(w, "Id da receita não informado", http.StatusBadRequest)
			return
		}
		receitaId, err := uuid.Parse(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		receita, err := re.ReceitaDB.FindByReceitaID(receitaId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		nomeDoRemedio := r.FormValue("nomeDoRemedio")
		receita.NomeDoRemedio = nomeDoRemedio
		err = re.ReceitaDB.UpdateReceita(receita)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func (re *ReceitaHandlers) DeleteReceita(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templ, err := template.ParseFiles("templates/receitas/deletar_receita.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		receitas, err := re.ReceitaDB.FindAllReceita()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Receitas []entity.Receita
		}{
			Receitas: receitas,
		}
		err = templ.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		id := r.FormValue("id")
		if id == "" {
			http.Error(w, "Id da receita não informado", http.StatusBadRequest)
			return
		}
		receitaId, err := uuid.Parse(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = re.ReceitaDB.DeleteReceita(receitaId)
		if err != nil {
			http.Error(w, "Erro ao deletar recepcionista", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

type MedicoHelper struct {
	MedicoDB database.MedicoInterface
}

func (re *ReceitaHandlers) verficarMedico(id uuid.UUID) bool {

	var db MedicoHelper
	medico := NewMedicoHandler(db.MedicoDB)
	return medico.FindByMedicoId(id)
}

type AgendamentoHelper struct {
	AgendamentoDB database.AgendamentoInterface
}

func (re *ReceitaHandlers) verficarAgendamento(id uuid.UUID) bool {

	var db AgendamentoHelper
	agendamento := NewAgendamentoHandler(db.AgendamentoDB)
	return agendamento.FindByAgendamentoId(id)
}

type PacienteHelper struct {
	Paciente database.PacienteInterface
}

func (re *ReceitaHandlers) verficarPaciente(id uuid.UUID) bool {
	var db PacienteHelper
	paciente := NewPacienteHandler(db.Paciente)
	return paciente.FindByPacienteId(id)
}
