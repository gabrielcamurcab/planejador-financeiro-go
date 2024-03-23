package main

import (
	"log"
	"os"

	"github.com/gabrielcamurcab/planejador-financeiro-go/adapter/http"
	"github.com/gabrielcamurcab/planejador-financeiro-go/db/postgres"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	os.Setenv("DOTENV", "../../.env")
	// Obtenha uma instância de *sql.DB do pacote db/postgres
	db, err := postgres.Connect()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	// Inicie o servidor HTTP passando a instância do banco de dados
	http.Init(db)
}
