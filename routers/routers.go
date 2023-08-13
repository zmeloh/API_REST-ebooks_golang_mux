package routers

import (
	"github.com/gorilla/mux"
	"example/api/handlers"
	"example/api/middleware"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Server listener 
	router.Use(middleware.LoggingMiddleware)
	
	// User router
	router.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/delete/{id}", handlers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/update/{id}", handlers.UpdateUser).Methods("PUT")

	// Category router
	router.HandleFunc("/categories", handlers.GetAllCategories).Methods("GET")
	router.HandleFunc("/categories/{id}", handlers.GetCategoryByID).Methods("GET")
	router.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")
	router.HandleFunc("/categories/delete/{id}", handlers.DeleteCategory).Methods("DELETE")
	router.HandleFunc("/categories/update/{id}", handlers.UpdateCategory).Methods("PUT")

	return router
}
