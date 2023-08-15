package models

type DatabaseInfo struct {
	DB_DRIVER string `json:"db_driver"`
	DB_USER string `json:"db_user"`
	DB_PASSWORD string `json:"db_password"`
	DB_NAME string `json:"db_name"`
	DB_HOST string `json:"db_host"`
	DB_PORT string `json:"db_port"`
		
}
