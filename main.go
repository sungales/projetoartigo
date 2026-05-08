package main

import (
	"log"
	"net/http"

	"github.com/sungales/projetoartigo/routes"

	database "github.com/sungales/projetoartigo/sql"
	_ "modernc.org/sqlite"
	
)

func main() {
	database.ConnectDatabase()

	http.HandleFunc("GET /artigos", routes.GetArticlesRoute)
	http.HandleFunc("POST /artigos/criar", routes.CreateArticleRoute)
	http.HandleFunc("GET /artigos/{id}", routes.GetArticleByIDRoute)


	log.Fatal(http.ListenAndServe(":8080", nil))
}
