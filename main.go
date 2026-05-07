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

	http.HandleFunc("/artigos", routes.GetArticlesRoute)
	http.HandleFunc("/artigos/criar", routes.CreateArticleRoute)
	http.HandleFunc("/artigos/{id}", routes.GetArticleByIDRoute)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
