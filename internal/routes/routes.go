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
	w.Header().Set("Content-Type", "text/html; charset=utf8")

	var artigo models.Artigo
	err := json.NewDecoder(r.Body).Decode(&artigo)
	if err != nil {
		http.Error(w, "erro ao decodificar o corpo da requisição", http.StatusBadRequest)
		return
	}

	artigo = models.Artigo{
		ID:        artigo.ID,
		Descricao: artigo.Descricao,
		CreatedAt: time.Now(),
	}

	if artigo.Descricao == "" {
		fmt.Print("a descricao nao pode estar vazia")
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

	component := templates.GetArtigosTemplate(artigos)
	if err = component.Render(r.Context(), w); err != nil {
		fmt.Println("erro ao renderizar")
		return
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
