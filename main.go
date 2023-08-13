package main

import (
	"fmt"
	"net/http"
	"example/api/routers"
	"example/api/database"
)

func main() {
	database.InitDB()

	router := routers.SetupRouter()

	fmt.Println("Server is running on :9090...")
	http.Handle("/", router)
	http.ListenAndServe(":9090", router)
}
