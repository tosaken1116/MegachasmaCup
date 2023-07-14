package config

import (
	"log"
	"os"

	"github.com/subosito/gotenv"
)
type appEnv struct {
	PostgresEnv *PostgresEnv
}
type PostgresEnv struct {
	DbHost    string
	DbUser    string
	DbPass    string
	DbName    string
	DbPort    string
}
func LoadEnv()*appEnv{
	if err := gotenv.Load(".env"); err != nil {
		log.Fatal("failed load env")
	}
	dbHost := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort := os.Getenv("POSTGRES_PORT")
	postgresEnv := &PostgresEnv{
		DbHost: dbHost,
		DbUser: dbUser,
		DbPass: dbPass,
		DbName: dbName,
		DbPort: dbPort,
	}
	conf := appEnv{
		PostgresEnv: postgresEnv,
	}
	return &conf
}