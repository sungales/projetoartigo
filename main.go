package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sungales/projetoartigo/models"
	_ "modernc.org/sqlite"
)

func main() {

	database, err := sql.Open("sqlite", "database.db")
	if err != nil {
		fmt.Println("Erro ao conectar com o banco de dados")
		return
	}

	fmt.Print("deu certo o banco")
	fmt.Print(database)

	articlesRoute := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Método não permitido", http.StatusBadRequest)
			return
		}

		w.Write([]byte("Nenhum artigo ainda"))
	}

	createArticle := func(w http.ResponseWriter, r *http.Request) {
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

		w.Write([]byte("Criado com sucesso"))

	}

	http.HandleFunc("/artigos", articlesRoute)
	http.HandleFunc("/artigos/criar", createArticle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
