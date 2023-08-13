package database

import (
    "database/sql"
    "log"
    "fmt"
    "os"
    _ "github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    dbDriver := os.Getenv("DB_DRIVER")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	DB, err = sql.Open(dbDriver, connectionString)
	if err != nil {
		log.Fatal(err)
	}
}