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
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusBadRequest)
		return
	}

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
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusBadRequest)
		return
	}

	var artigos = database.GetAllArticles()
	artigosJson, err := json.Marshal(artigos)
	if err != nil {
		fmt.Println("erro ao converter para JSON: ", err)
	}
	w.Write(artigosJson)
	fmt.Println("artigos enviados")
}

func GetArticleByIDRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "método inválido", http.StatusBadRequest)
		return
	}

	var artigoID = r.PathValue("id")
	id, err := strconv.Atoi(artigoID)
	if err != nil {
		fmt.Println("não foi possivel converter: ", err)
	}

	artigo := database.GetArticleByID(id)
	artigoJSON, err := json.Marshal(artigo)
	if err != nil {
		fmt.Println("não foi possivel transcrever para JSON: ", err)
	}
	w.Write(artigoJSON)
	fmt.Println("artigo enviado")
}
