package main

import (
	"fmt"
	"net/http"
	"example/api/routers"
	"example/api/dao"
)

func main() {
	dao.InitDB()

	router := routers.SetupRouter()

	fmt.Println("Server is running on :9090...")
	http.Handle("/", router)
	http.ListenAndServe(":9090", router)
}
