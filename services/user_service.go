package services

import (
	"example/api/controllers"
	"example/api/models"
)

// GetAllUsers récupère tous les utilisateurs.
func GetAllUsers() (u []models.User) {
	u, _ = controllers.GetAllUsers()
	return
}

// GetUserByID récupère un utilisateur par son ID.
func GetUserByID(id int) (u models.User) {
	u, _ = controllers.GetUserByID(id)
	return
}

// InsertUser insère un nouvel utilisateur.
func InsertUser(u *models.User) error {
	return controllers.InsertUser(u)

}

// UpdateUser met à jour un utilisateur existant.
func UpdateUser(id int, u *models.User) error {
	err := controllers.UpdateUser(id, u)
	return err
}

// DeleteUser supprime un utilisateur par son ID.
func DeleteUser(id int) error {
	err := controllers.DeleteUser(id)
	return err
}
