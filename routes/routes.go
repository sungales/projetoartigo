package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sungales/projetoartigo/models"
	database "github.com/sungales/projetoartigo/sql"
)

func CreateArticleRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusBadRequest)
		return
	}

	var novoArtigo models.Artigo

	// pega o corpo da requisição e passa pra variável novoArtigo
	err := json.NewDecoder(r.Body).Decode(&novoArtigo)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da requisição")
		return
	}
	fmt.Println("artigo criado com sucesso!")
}

func GetArticlesRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusBadRequest)
		return
	}

	// VERIFICAR SE AS ROTAS ESTÃO FUNCIONANDO
	// TERMINAR DE ORGANIZAR O CÓDIGO DE SQL/LÓGICA DO BANCO
	// COLOCAR TUDO CERTO NA main.go

	var artigos = database.GetArticles()
	fmt.Println(artigos)
}
