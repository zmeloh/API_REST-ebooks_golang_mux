package dao

import (
	"database/sql"
	"encoding/json"
	"example/api/models"
	"fmt"
	"log"
	"os"

	//"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	// Lecture du fichier JSON
	data,err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
	}

	var databaseinfo models.Database

	if err = json.Unmarshal(data, &databaseinfo); err != nil {
		fmt.Println(err)
	}


	dbDriver := databaseinfo.DB_DRIVER //os.Getenv("DB_DRIVER")
	dbUser := databaseinfo.DB_USER //os.Getenv("DB_USER")
	dbPassword := databaseinfo.DB_PASSWORD //os.Getenv("DB_PASSWORD")
	dbHost := databaseinfo.DB_HOST //os.Getenv("DB_HOST")
	dbPort := databaseinfo.DB_PORT //os.Getenv("DB_PORT")
	dbName := databaseinfo.DB_NAME //os.Getenv("DB_NAME")

	// Fermeture du fichier JSON
	

	connectionString := ""

	switch dbDriver {
	case "mysql":
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	case "sqlite":
		connectionString = dbName + ".db"
	case "postgres":
		connectionString = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbUser, dbPassword, dbName, dbHost, dbPort)
	default:
		log.Fatal("Unsupported database driver")
	}

	DB, err = sql.Open(dbDriver, connectionString)
	if err != nil {
		log.Fatal(err)
	}
}
