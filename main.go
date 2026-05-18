package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/sungales/projetoartigo/internal/routes"
	"github.com/sungales/projetoartigo/middleware"

	database "github.com/sungales/projetoartigo/db"
	_ "modernc.org/sqlite"
)

func main() {
	if err := database.ConnectDatabase(); err != nil {
		log.Fatal("não foi possivel conectar ao banco: ", err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("não foi possivel carregar o .env ", err)
		return
	}

	http.HandleFunc("GET /artigos", routes.GetAllArticlesRoute)
	http.HandleFunc("/artigos/criar", routes.CreateArticleRoute)
	http.HandleFunc("GET /artigos/id/{id}", routes.GetArticleByIDRoute)
	http.Handle("/artigos/editar/{id}", middleware.APIKeyAuth(http.HandlerFunc(routes.EditArticleRoute)))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
