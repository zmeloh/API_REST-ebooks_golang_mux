package services

import (
	"example/api/models"
)

// GetAllUsers récupère tous les utilisateurs.
func GetAllUsers(){
	// controllers.GetAllUsers()
}

// GetUserByID récupère un utilisateur par son ID.
func GetUserByID(id int){
	// controllers.GetUserByID(id)
}

// InsertUser insère un nouvel utilisateur.
func InsertUser(u models.User){
	// controllers.InsertUser(u)
}

// UpdateUser met à jour un utilisateur existant.
func UpdateUser(id int, u models.User){
	// controllers.UpdateUser(id, u)
}

// DeleteUser supprime un utilisateur par son ID.
func DeleteUser(id int){
	// controllers.DeleteUser(id)
}
