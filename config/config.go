package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type DBConfig struct {
	Db *sql.DB
}

func NewDBConfig() *DBConfig {
	var (
		dbName     string
		dbHost     string
		dbPort     string
		dbUsername string
		dbPassword string
	)
	dbName = os.Getenv("DB_DATABASE")
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUsername = os.Getenv("DB_USERNAME")
	dbPassword = os.Getenv("DB_PASSWORD")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUsername, dbPassword, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
	}

	return &DBConfig{Db: db}
}
