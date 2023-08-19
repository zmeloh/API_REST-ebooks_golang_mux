package routers

import (
	"github.com/gorilla/mux"
	"example/api/rest"
	"example/api/middleware"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Serveur listener 
	router.Use(middleware.LoggingMiddleware)
	
	// Users routers
	router.HandleFunc("/users", rest.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", rest.GetUserByID).Methods("GET")
	router.HandleFunc("/users", rest.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", rest.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}", rest.UpdateUser).Methods("PUT")

	// Categories routers
	router.HandleFunc("/categories", rest.GetAllCategories).Methods("GET")
	router.HandleFunc("/categories/{id}", rest.GetCategoryByID).Methods("GET")
	router.HandleFunc("/categories", rest.CreateCategory).Methods("POST")
	router.HandleFunc("/categories/{id}", rest.DeleteCategory).Methods("DELETE")
	router.HandleFunc("/categories/{id}", rest.UpdateCategory).Methods("PUT")
	

	// Ebooks routers
	router.HandleFunc("/ebooks", rest.GetAllEbooks).Methods("GET")
	router.HandleFunc("/ebooks", rest.CreateEbook).Methods("POST")
	router.HandleFunc("/ebooks/{id}", rest.GetEbookByID).Methods("GET")
	router.HandleFunc("/ebooks/category/{id}", rest.GetEbookByCategoryID).Methods("GET")
	router.HandleFunc("/ebooks/{id}", rest.DeleteEbook).Methods("DELETE")
	router.HandleFunc("/ebooks/{id}", rest.UpdateEbook).Methods("PUT")

	
	// Favorites routers
	router.HandleFunc("/favorites", rest.GetAllFavorites).Methods("GET")
	router.HandleFunc("/favorites/{id}", rest.GetFavoriteByID).Methods("GET")
	router.HandleFunc("/favorites/user/{id}", rest.GetFavoritesByUserID).Methods("GET")
	router.HandleFunc("/favorites/ebook/{id}", rest.GetFavoritesByEbookID).Methods("GET")
	router.HandleFunc("/favorites", rest.CreateFavorite).Methods("POST")
	router.HandleFunc("/favorites/id}", rest.DeleteFavorite).Methods("DELETE")
	router.HandleFunc("/favorites/id}", rest.UpdateFavorite).Methods("PUT")
	
	return router
}
