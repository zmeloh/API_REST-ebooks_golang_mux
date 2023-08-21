package controllers

import (
	"example/api/dao"
	"example/api/models"
	"example/api/utils"
	"fmt"
)

func InsertUser(u *models.User) error {

	result, _ := dao.ExistingUser(u.Email, u.Username)
	fmt.Print(result)
	if u.Email == result.Email {
		return fmt.Errorf("email already exist")
	}

	if u.Username == result.Username {
		return fmt.Errorf("username already exist")
	}

	err := dao.InsertUser(u)
	if err != nil {
		utils.Logger(err)
		return err
	}

	return nil
}

func GetAllUsers() ([]models.User, error) {
	users, err := dao.SelectAllUsers()
	if err != nil {
		utils.Logger(err)
		return nil, err
	}
	return users, nil
}

func GetUserByID(id int) (models.User, error) {
	user, err := dao.SelectUserByID(id)
	if err != nil {
		utils.Logger(err)
		return models.User{}, err
	}
	return user, nil
}

func ExistingUser(email, username string) (models.User, error) {
	user, err := dao.ExistingUser(email, username)
	if err != nil {
		utils.Logger(err)
		return models.User{}, err
	}
	return user, nil
}

func UpdateUser(id int, updatedUser *models.User) error {

	existingUser, err := dao.SelectUserByID(id)
	if err != nil {
		utils.Logger(err)
		return err
	}
	existingUser.Username = updatedUser.Username
	existingUser.Email = updatedUser.Email
	updatedUser.ID = existingUser.ID
	_, err = dao.UpdateUser(existingUser)
	if err != nil {
		utils.Logger(err)
		return err
	}
	return err
}

func DeleteUser(id int) error {
	err := dao.DeleteUser(id)
	if err != nil {
		utils.Logger(err)
		return err
	}
	return err
}
