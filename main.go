package main

import (
	"log"
	"net/http"

	"github.com/sungales/projetoartigo/routes"

	database "github.com/sungales/projetoartigo/sql"
	_ "modernc.org/sqlite"
)

func main() {
	err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("não foi possivel conectar ao banco: ", err)
	}

	http.HandleFunc("GET /artigos", routes.GetArticlesRoute)
	http.HandleFunc("POST /artigos/criar", routes.CreateArticleRoute)
	http.HandleFunc("GET /artigos/{id}", routes.GetArticleByIDRoute)

	// VALIDAR TIMESTAMP PARA O CREATED_AT
	// TENTAR ORGANIZAR E ESTILIZAR PELO MENOS O GET DOS ARTIGOS

	log.Fatal(http.ListenAndServe(":8080", nil))
}
