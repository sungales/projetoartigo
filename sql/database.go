package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/sungales/projetoartigo/models"
)

var database *sql.DB

func ConnectDatabase() {
	var err error

	database, err = sql.Open("sqlite", "./sql/database.db")
	if err != nil {
		log.Fatal("não foi possivel conectar ao banco ", err)
	}

	fmt.Println("banco funcionando")

	sqlfile, err := os.ReadFile("./sql/sql-manager.sql")
	if err != nil {
		fmt.Println("não foi possivel ler o arquivo sql: ", err)
	}

	_, err = database.Exec(string(sqlfile))
	if err != nil {
		fmt.Println("não foi possível criar as tabelas no banco: ", err)
	}
}

func GetAllArticles() []models.Artigo {
	query := "SELECT * FROM artigos"

	rows, err := database.Query(query)
	if err != nil {
		fmt.Println("não foi possivel trazer os artigos do banco: ", err)
	}
	defer rows.Close()

	var artigos []models.Artigo

	for rows.Next() {
		var artigo models.Artigo
		err := rows.Scan(&artigo.ID, &artigo.Descricao, &artigo.CreatedAt)
		if err != nil {
			fmt.Println("não foi possivel ler os artigos do banco: ", err)
		}
		artigos = append(artigos, artigo)
	}
	return artigos
}

func GetArticleByID(id int) models.Artigo {
	query := "SELECT * FROM artigos WHERE id = ?"

	var artigo models.Artigo
	err := database.QueryRow(query, id).Scan(&artigo.ID, &artigo.Descricao, &artigo.CreatedAt)
	if err != nil {
		fmt.Println("não foi possivel trazer o artigo", err)
	}
	return artigo
}

func CreateArticle(artigo models.Artigo) { 
	query := "INSERT INTO artigos (descricao, created_at) VALUES (?, ?)"
    
	_, err := database.Exec(query, artigo.Descricao, artigo.CreatedAt)
	if err != nil { 
		fmt.Println("não foi possivel criar o artigo ", err)
	}
	
	fmt.Println("artigo criado!")
}