package dao

import (
	"database/sql"
	"encoding/json"
	"example/api/models"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	// Lecture du fichier JSON
	data, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal("Erreur lors de la lecture du fichier de configuration:", err)
	}

	var config models.Config

	if err = json.Unmarshal(data, &config); err != nil {
		log.Fatal("Erreur lors du décodage de la configuration JSON:", err)
	}

	// Utilisation des valeurs de configuration pour la connexion à la base de données
	dbDriver := config.DatabaseInfo.DBDriver
	dbUser := config.DatabaseInfo.DBUser
	dbPassword := config.DatabaseInfo.DBPassword
	dbHost := config.DatabaseInfo.DBHost
	dbPort := config.DatabaseInfo.DBPort
	dbName := config.DatabaseInfo.DBName
	connectionString := ""

	fmt.Println(string(data))

	// if dbDriver == "mysql" {
	// 	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	// } else if dbDriver == "postgresql" {
	// 	connectionString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	// } else {
	// 	log.Fatal("Pilote de base de données non pris en charge")

	// }

	switch dbDriver {
	case "mysql":
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	case "postgres":

		connectionString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	default:
		log.Fatal("Pilote de base de données non pris en charge")
	}

	DB, err = sql.Open(dbDriver, connectionString)
	if err != nil {
		log.Fatal("Erreur lors de la connexion à la base de données:", err)
		return
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Erreur:", err)
	}
}
