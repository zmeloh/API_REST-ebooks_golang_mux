package dao

import (
	"database/sql"
	"encoding/json"
	"example/api/models"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	// Lecture du fichier JSON
	data, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal("Erreur lors de la lecture du fichier de configuration:", err)
	}

	var databaseInfo models.DatabaseInfo

	if err = json.Unmarshal(data, &databaseInfo); err != nil {
		log.Fatal("Erreur lors du décodage de la configuration JSON:", err)
	}

	// Utilisation des valeurs de configuration pour la connexion à la base de données
	dbDriver := databaseInfo.DBDriver
	dbUser := databaseInfo.DBUser
	dbPassword := databaseInfo.DBPassword
	dbHost := databaseInfo.DBHost
	dbPort := databaseInfo.DBPort
	dbName := databaseInfo.DBName

	connectionString := ""

	switch dbDriver {
	case "mysql":
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	default:
		log.Fatal("Pilote de base de données non pris en charge")
	}

	DB, err = sql.Open(dbDriver, connectionString)
	if err != nil {
		log.Fatal("Erreur lors de la connexion à la base de données:", err)
	}
}
