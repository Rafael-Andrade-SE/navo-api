package database

import (
	"fmt"
	"os"

	"github.com/rafael-andrade-se/navo/api/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConectarPostgres() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Verificação básica de variáveis
	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		logger.Log.Error("Variáveis de ambiente do banco estão incompletas")
		return nil, fmt.Errorf("variáveis de ambiente do banco estão incompletas")
	}

	// Montar DSN
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	logger.Log.Info("Conectando ao banco PostgreSQL...")

	// Abrir conexão
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Error("Erro ao conectar no banco de dados", "erro", err)
		return nil, err
	}

	logger.Log.Info("Conexão com banco de dados estabelecida com sucesso ✅")
	return db, nil
}
