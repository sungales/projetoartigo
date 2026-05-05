package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sungales/projetoartigo/models"
)

func main() {

	var database = []models.Artigo{}

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

		// arrumar a logica de ID, por enquanto é só o tamanho do banco de dados + 1
		// arrumar a logica do artigo em si, por enquanto é apenas o que vem na requisição.
		fmt.Print(novoArtigo)
		
		fmt.Print(database)
	}

	http.HandleFunc("/artigos", articlesRoute)
	http.HandleFunc("/artigos/criar", createArticle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
