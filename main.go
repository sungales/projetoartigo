package main

import (
	"context"
	"log"
	"net/http"

	"github.com/sungales/projetoartigo/html/templates"
	"github.com/sungales/projetoartigo/internal/routes"

	database "github.com/sungales/projetoartigo/db"
	_ "modernc.org/sqlite"
)

func main() {
	err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("não foi possivel conectar ao banco: ", err)
	}

	http.HandleFunc("GET /artigos", routes.GetArticlesRoute, ) 
	http.HandleFunc("POST /artigos/criar", routes.CreateArticleRoute)
	http.HandleFunc("GET /artigos/{id}", routes.GetArticleByIDRoute)
	http.HandleFunc("/teste/teste", func(w http.ResponseWriter, r *http.Request) {
		component := templates.CriarArtigoTemplate()
		component.Render(context.Background(), w)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
