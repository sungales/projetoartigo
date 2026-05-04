package main

import (
	"log"
	"net/http"
	"time"
)

func main() { 

	articleRoute := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Método não permitido", http.StatusBadRequest)
			return
		}

		w.Write([]byte("Nenhum artigo ainda"))
	}

	createArticleRoute := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusBadRequest)
			return
		}

		Artigo := &Artigo{
			ID: 1,
			Descricao: "Na epoca da fruta, todo mundo era banana",
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		}


		w.Write([]byte(Artigo.Descricao))
	}


	http.HandleFunc("/artigos", articleRoute)
	http.HandleFunc("/artigos/criar", createArticleRoute)

	log.Fatal(http.ListenAndServe(":8080", nil))
}