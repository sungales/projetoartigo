package main

import (
	"log"
	"net/http"

	"github.com/sungales/projetoartigo/routes"
	_ "modernc.org/sqlite"
)

func main() {

	http.HandleFunc("/artigos", routes.GetArticlesRoute)
	http.HandleFunc("/artigos/criar", routes.CreateArticleRoute)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
