package models

// DatabaseInfo contient les informations de configuration de la base de données.
type DatabaseInfo struct {
	DBDriver   string `json:"db_driver"`
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
	DBName     string `json:"db_name"`
	DBHost     string `json:"db_host"`
	DBPort     string `json:"db_port"`
}

type Config struct {
	DatabaseInfo DatabaseInfo `json:"database_info"`
}
