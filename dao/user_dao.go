package dao

import (
	"database/sql"
	"example/api/models"
	"example/api/utils"
	"fmt"
)

// CreateUser crée un nouvel utilisateur en utilisant les données du corps de la requête.
func InsertUser(u models.User) error {
	// Exécute la requête pour insérer un nouvel utilisateur dans la base de données
	result, err := DB.Exec("INSERT INTO users (username, email) VALUES (?, ?)", u.Username, u.Email)
	if err != nil {
		return err
	}
	// Obtient l'ID généré lors de l'insertion
	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Met à jour l'ID dans l'objet User
	u.ID = int(userID)
	utils.Logger()
	return err
}

func SelectAllUsers()([]models.User, error) {

	var users []models.User

	// Exécute la requête pour récupérer les informations de l'utilisateur depuis la base de données
	rows, err := DB.Query("SELECT id, username, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func SelectUserByID(id int)(models.User, error) {
	var user models.User
	// Exécute la requête pour récupérer les informations de l'utilisateur depuis la base de données
	err := DB.QueryRow("SELECT id, username, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("No user found with ID %d", id)
		}
		return models.User{}, err
	}
	return user, nil
}

func UpdateUser(id int, updatedUser models.User) (models.User, error){

	// Requête pour mettre à jour les informations du livre électronique dans la base de données
	_, err := DB.Exec("UPDATE users SET username = ?, email= ? WHERE id = ?", updatedUser.Username, updatedUser.Email, id)
	if err != nil {
		return models.User{}, err
	}

	// Mettre à jour l'ID dans l'objet Ebook
	updatedUser.ID = id

	return updatedUser, nil
}

func DeleteUser(id int) error{
	// Requête pour supprimer un livre électronique par ID dans la base de données
	_, err := DB.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
