package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sungales/projetoartigo/models"
	database "github.com/sungales/projetoartigo/sql"
	"github.com/sungales/projetoartigo/templates"
)

func CreateArticleRoute(w http.ResponseWriter, r *http.Request) {
	var artigo models.Artigo

	err := json.NewDecoder(r.Body).Decode(&artigo)
	if err != nil {
		http.Error(w, "erro ao decodificar o corpo da requisição", http.StatusBadRequest)
		return
	}

	if artigo.Descricao == "" || artigo.CreatedAt == "" {
		http.Error(w, "falta preencher a descricao", http.StatusBadRequest)
		return
	}

	database.CreateArticle(artigo)
	w.Write([]byte("artigo criado!"))
}

func GetArticlesRoute(w http.ResponseWriter, r *http.Request) {
	var artigos, err = database.GetAllArticles()
	if err != nil {
		fmt.Println("não foi possivel pegar os artigos do banco: ", err)
		return
	}

	component := templates.ArtigoTemplate(artigos)
	if err = component.Render(r.Context(), w); err != nil {
		fmt.Println("erro ao renderizar")
		return
	}
}

func GetArticleByIDRoute(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println("artigo com o ID ", id, " enviado")
}
