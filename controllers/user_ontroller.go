package controllers

import (
	"example/api/dao"
	"example/api/models"
)

func InsertUser(u models.User) error {
	err := dao.InsertUser(u)
	if err != nil {
		return err
	}
	return nil
}

func GetAllUsers() ([]models.User, error) {
	users, err := dao.SelectAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(id int) (models.User, error) {
	user, err := dao.SelectUserByID(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func UpdateUser(id int, updatedUser models.User) error {
	existingUser, err := dao.SelectUserByID(id)
	if err != nil {
		return err
	}
	existingUser.Username = updatedUser.Username
	existingUser.Email = updatedUser.Email
	_, err = dao.UpdateUser(id, updatedUser)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	err := dao.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
