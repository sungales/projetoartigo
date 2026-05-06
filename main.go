package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sungales/projetoartigo/models"
	_ "modernc.org/sqlite"
)

func main() {

	database, err := sql.Open("sqlite", "database.db")
	if err != nil {
		fmt.Println("Erro ao conectar com o banco de dados")
		return
	}
	defer database.Close()

	fmt.Println("banco funcionando!")

	sql, err := os.ReadFile("./models/create-table.sql")
	if err != nil {
		fmt.Println("erro ao ler o arquivo SQL")
	}

	_, err = database.Exec(string(sql))
	if err != nil {
		log.Fatal("erro ao criar a tabela: ", err)
	}

	articlesRoute := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Método não permitido", http.StatusBadRequest)
			return
		}

		rows, err := database.Query("SELECT * FROM artigos")
		if err != nil {
			fmt.Println("erro ao puxar os artigos do banco: ", err)
		}
		defer rows.Close()
		for rows.Next() {
			var artigo models.Artigo

			err := rows.Scan(&artigo.ID, &artigo.Descricao, &artigo.CreatedAt)
			if err != nil {
				fmt.Println("não foi possivel ler os artigos do banco: ", err)
			}
			res, err := json.Marshal(artigo)
			if err != nil {
				fmt.Println("erro ao codificar o artigo para JSON: ", err)
			}
			w.Write([]byte("Todos os artigos: \n"))
			w.Write(res)

		}
	}

	createArticle := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusBadRequest)
			return
		}

		var novoArtigo models.Artigo
		var template = "INSERT INTO artigos (descricao, created_at) VALUES (?, ?)"

		// pega o corpo da requisição e passa pra variável novoArtigo
		err := json.NewDecoder(r.Body).Decode(&novoArtigo)
		if err != nil {
			fmt.Println("Erro ao ler o corpo da requisição")
			return
		}

		_, err = database.Exec(template, novoArtigo.Descricao, novoArtigo.CreatedAt)
		if err != nil {
			fmt.Println("erro ao inserir o artigo no banco: ", err)
			return
		}
		fmt.Println("artigo criado com sucesso!")

	}

	http.HandleFunc("/artigos", articlesRoute)
	http.HandleFunc("/artigos/criar", createArticle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
