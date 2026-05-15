package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/sungales/projetoartigo/internal/models"
)

var database *sql.DB

func ConnectDatabase() error {
	var err error

	database, err = sql.Open("sqlite", "./database.db")
	if err != nil {
		log.Fatal("não foi possivel conectar ao banco ", err)
		return err
	}

	fmt.Println("banco funcionando")

	sqlfile, err := os.ReadFile("./db/sql-manager.sql")
	if err != nil {
		fmt.Println("não foi possivel ler o arquivo sql: ", err)
		return err
	}

	_, err = database.Exec(string(sqlfile))
	if err != nil {
		fmt.Println("não foi possível criar as tabelas no banco: ", err)
		return err
	}
	return nil
}

func GetAllArticles() ([]models.Artigo, error) {
	query := "SELECT id, titulo, descricao, created_at FROM artigos"

	rows, err := database.Query(query)
	if err != nil {
		fmt.Println("não foi possivel trazer os artigos do banco: ", err)
		return []models.Artigo{}, err
	}
	defer rows.Close()

	var artigos []models.Artigo

	for rows.Next() {
		var artigo models.Artigo
		err := rows.Scan(&artigo.ID, &artigo.Titulo, &artigo.Descricao, &artigo.CreatedAt)
		if err != nil {
			fmt.Println("não foi possivel ler os artigos do banco: ", err)
			return []models.Artigo{}, err
		}
		artigos = append(artigos, artigo)
	}
	return artigos, nil
}

func GetArticleByID(id int) (models.Artigo, error) {
	query := "SELECT id, titulo, descricao, created_at FROM artigos WHERE id = ?"

	var artigo models.Artigo
	err := database.QueryRow(query, id).Scan(&artigo.ID, &artigo.Titulo, &artigo.Descricao, &artigo.CreatedAt)
	if err != nil {
		fmt.Println("não foi possivel trazer o artigo", err)
		return models.Artigo{}, err
	}
	return artigo, nil
}

func CreateArticle(artigo models.Artigo) error {
	query := "INSERT INTO artigos (titulo, descricao, created_at) VALUES (?, ?, ?)"

	_, err := database.Exec(query, artigo.Titulo, artigo.Descricao, artigo.CreatedAt)
	if err != nil {
		fmt.Println("não foi possivel criar o artigo ", err)
		return err
	}

	fmt.Println("artigo criado!")
	return nil
}
