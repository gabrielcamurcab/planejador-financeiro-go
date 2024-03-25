// db/postgres/main_test.go

package postgres

import (
	"os"
	"testing"
)

func TestConnect(t *testing.T) {
	tempHost := os.Getenv("DB_HOST")
	tempPort := os.Getenv("DB_PORT")
	tempUser := os.Getenv("DB_USER")
	tempPassword := os.Getenv("DB_PASSWORD")
	tempDBName := os.Getenv("DB_NAME")

	os.Setenv("DB_HOST", "postgres")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "root")
	os.Setenv("DB_NAME", "go_test")

	_, err := Connect()
	if err != nil {
		t.Errorf("Erro ao conectar: %v", err)
	}

	os.Setenv("DB_HOST", tempHost)
	os.Setenv("DB_PORT", tempPort)
	os.Setenv("DB_USER", tempUser)
	os.Setenv("DB_PASSWORD", tempPassword)
	os.Setenv("DB_NAME", tempDBName)
}
