package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/sungales/projetoartigo/html/templates"
	"github.com/sungales/projetoartigo/internal/models"

	database "github.com/sungales/projetoartigo/db"
)

func CreateArticleRoute(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html; charset=utf8")

		component := templates.CriarArtigoTemplate()
		if err := component.Render(r.Context(), w); err != nil {
			http.Error(w, "erro ao tentar renderizar o Templ", http.StatusInternalServerError)
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "erro ao tentar ler tudo do body", http.StatusInternalServerError)
			return
		}
	}

	if r.Method == http.MethodPost {
		artigoTexto := r.FormValue("textoDoArtigo")
		artigoTitulo := r.FormValue("tituloDoArtigo")

		var artigo models.Artigo

		fmt.Println("Descricao do artigo: ", artigoTexto, " Titulo do artigo: ")

		artigo = models.Artigo{
			ID:        artigo.ID,
			Titulo:    artigoTitulo,
			Descricao: artigoTexto,
			CreatedAt: time.Now(),
		}

		if artigoTexto == "" || artigoTitulo == "" {
			fmt.Print("a descricao e o titulo nao podem estar vazios")
			return
		}

		if err := database.CreateArticle(artigo); err != nil {
			fmt.Println("erro ao criar artigo")
			return
		}
	}
}

func GetAllArticlesRoute(w http.ResponseWriter, r *http.Request) {
	var artigos, err = database.GetAllArticles()
	if err != nil {
		fmt.Println("não foi possivel pegar os artigos do banco: ", err)
		return
	}

	component := templates.GetAllArtigosTemplate(artigos)
	if err = component.Render(r.Context(), w); err != nil {
		fmt.Println("erro ao renderizar")
		return
	}
}

func EditArticleRoute(w http.ResponseWriter, r *http.Request) {
	var err error
	var idstr = r.PathValue("id")
	var id int

	if idstr == "" {
		http.Error(w, "você não especificou um ID", http.StatusBadRequest)
		return
	}

	id, err = strconv.Atoi(idstr)
	if err != nil {
		fmt.Println("não foi possivel converter o ID para número: ", err)
		return
	}

	if artigo, err := database.GetArticleByID(id); err != nil {
		fmt.Println("erro ao tentar pegar o artigo pelo ID", err)
		return
	} else {
		if err = json.NewDecoder(r.Body).Decode(&artigo); err != nil {
			fmt.Print("erro ao puxar/ler o body da requisição ", err)
			return
		}

		if err = database.EditArticle(id, artigo); err != nil {
			fmt.Println("erro ao editar o artigo: ", err)
			return
		}
		fmt.Println("arquivo editado com sucesso!")
	}
}

func GetArticleByIDRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")

	var artigoID = r.PathValue("id")

	id, err := strconv.Atoi(artigoID)
	if err != nil {
		fmt.Println("não foi possivel converter: ", err)
		http.Error(w, "ID inválido, precisa ser um número", http.StatusBadRequest)
		return
	}

	artigo, err := database.GetArticleByID(id)
	if err != nil {
		fmt.Println("não foi possivel pegar o artigo do banco: ", err)
		return
	}

	artigoJSON, err := json.Marshal(artigo)
	if err != nil {
		fmt.Println("não foi possivel transcrever para JSON: ", err)
		return
	}
	w.Write(artigoJSON)
	fmt.Println("artigo com o ID: ", id, " enviado")
}
