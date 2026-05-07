package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sungales/projetoartigo/models"
	database "github.com/sungales/projetoartigo/sql"
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
	var artigos = database.GetAllArticles()

	artigosJson, err := json.Marshal(artigos)
	if err != nil {
		fmt.Println("erro ao converter para JSON: ", err)
		return
	}
	w.Write(artigosJson)
	fmt.Println("todos os artigos enviados!")
}

func GetArticleByIDRoute(w http.ResponseWriter, r *http.Request) {
	var artigoID = r.PathValue("id")

	id, err := strconv.Atoi(artigoID)
	if err != nil {
		fmt.Println("não foi possivel converter: ", err)
		http.Error(w, "ID inválido, precisa ser um número", http.StatusBadRequest)
		return
	}

	artigo := database.GetArticleByID(id)
	artigoJSON, err := json.Marshal(artigo)
	if err != nil {
		fmt.Println("não foi possivel transcrever para JSON: ", err)
		return
	}
	w.Write(artigoJSON)
	fmt.Println("artigo com o ID ", id, " enviado")
}
