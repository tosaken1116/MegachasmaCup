package db

import (
	"fmt"
	"log"
	"megachasma/config"
	model "megachasma/graph/model/db"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConnector struct {
	Conn *gorm.DB
}

func NewPostgresConnector() *PostgresConnector {
	conf := config.LoadEnv()
	dsn := postgresConnInfo(*conf.PostgresEnv)
	Psql, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return &PostgresConnector{
		Conn: Psql,
	}
}
func postgresConnInfo(postgresInfo config.PostgresEnv)string{
	databaseSourceName := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		postgresInfo.DbHost,
		postgresInfo.DbUser,
		postgresInfo.DbPass,
		postgresInfo.DbName,
		postgresInfo.DbPort,
	)
	return databaseSourceName
}

func MigrateDatabase()(){
	conf := config.LoadEnv()
	dsn := postgresConnInfo(*conf.PostgresEnv)
	Psql, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}
	if err :=Psql.AutoMigrate(&model.School{},&model.Class{},&model.User{},&model.Note{},&model.Tag{},&model.Comment{});err != nil{
		log.Fatal(err)
	}
}