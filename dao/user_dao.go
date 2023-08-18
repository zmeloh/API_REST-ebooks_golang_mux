package dao

import (
	"database/sql"
	"example/api/models"
	"example/api/utils"
	"fmt"
)

// CreateUser crée un nouvel utilisateur en utilisant les données du corps de la requête.
func InsertUser(u *models.User) error {
	// Exécute la requête pour insérer un nouvel utilisateur dans la base de données
	// result, err := DB.QueryRow(`INSERT INTO users (username, email) VALUES ($1, $2)`).Scan(&u.Username, &u.Email)
	err := DB.QueryRow("INSERT INTO users (username, email) VALUES ($1, $2) RETURNING 1", u.Username, u.Email).Scan(&u.ID)
	if err != nil {
		utils.Logger(err)
		return err
	}

	return err
}

func SelectAllUsers() ([]models.User, error) {

	var users []models.User

	// Exécute la requête pour récupérer les informations de l'utilisateur depuis la base de données
	rows, err := DB.Query("SELECT id, username, email FROM users ORDER BY id")
	if err != nil {
		utils.Logger(err)
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

func SelectUserByID(id int) (models.User, error) {
	var user models.User
	// Exécute la requête pour récupérer les informations de l'utilisateur depuis la base de données
	err := DB.QueryRow("SELECT id, username, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		utils.Logger(err)
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("no user found with ID %d", id)
		}
		return models.User{}, err
	}
	return user, nil
}

func UpdateUser(updatedUser models.User) (models.User, error) {

	// Requête pour mettre à jour les informations du livre électronique dans la base de données
	_, err := DB.Exec("UPDATE users SET username = $1, email= $2 WHERE id = $3", updatedUser.Username, updatedUser.Email, updatedUser.ID)
	if err != nil {
		utils.Logger(err)
		return models.User{}, err
	}

	return updatedUser, nil
}

func DeleteUser(id int) error {
	// Requête pour supprimer un livre électronique par ID dans la base de données
	_, err := DB.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		utils.Logger(err)
	}

	return err
}
