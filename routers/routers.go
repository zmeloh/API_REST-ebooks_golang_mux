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
	
	// Users routers
	router.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/delete/{id}", handlers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/update/{id}", handlers.UpdateUser).Methods("PUT")

	// Categories routers
	router.HandleFunc("/categories", handlers.GetAllCategories).Methods("GET")
	router.HandleFunc("/categories/{id}", handlers.GetCategoryByID).Methods("GET")
	router.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")
	router.HandleFunc("/categories/delete/{id}", handlers.DeleteCategory).Methods("DELETE")
	router.HandleFunc("/categories/update/{id}", handlers.UpdateCategory).Methods("PUT")


	// Ebooks routers
	router.HandleFunc("/ebooks", handlers.GetAllEbooks).Methods("GET")
	router.HandleFunc("/ebooks/{id}", handlers.GetEbookByID).Methods("GET")
	router.HandleFunc("/ebooks/category/{id}", handlers.GetEbooksByCategory).Methods("GET")
	router.HandleFunc("/ebooks", handlers.CreateEbook).Methods("POST")
	router.HandleFunc("/ebooks/delete/{id}", handlers.DeleteEbook).Methods("DELETE")
	router.HandleFunc("/ebooks/update/{id}", handlers.UpdateEbook).Methods("PUT")

	// Favorites routers
	router.HandleFunc("/favorites", handlers.GetAllFavorites).Methods("GET")
	router.HandleFunc("/favorites/{id}", handlers.GetFavoriteByID).Methods("GET")
	router.HandleFunc("/favorites/user/{id}", handlers.GetFavoritesByUserID).Methods("GET")
	router.HandleFunc("/favorites/ebook/{id}", handlers.GetFavoritesByEbookID).Methods("GET")
	router.HandleFunc("/favorites", handlers.CreateFavorite).Methods("POST")
	router.HandleFunc("/favorites/delete/{id}", handlers.DeleteFavorite).Methods("DELETE")
	router.HandleFunc("/favorites/update/{id}", handlers.UpdateFavorite).Methods("PUT")

	return router
}
