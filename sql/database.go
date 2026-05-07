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

	sqlfile, err := os.ReadFile("./sql/sql-manager.sql")
	if err != nil {
		fmt.Println("não foi possivel ler o arquivo sql: ", err)
	}

	_, err = database.Exec(string(sqlfile))
	if err != nil {
		fmt.Println("não foi possível criar as tabelas no banco: ", err)
	}

}

func GetArticles() []models.Artigo {
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
