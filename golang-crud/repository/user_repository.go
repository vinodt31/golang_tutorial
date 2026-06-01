package repository

import (
	"errors"
	"strconv"

	"golang-crud/config"
	"golang-crud/model"
)

func CreateUser(user *model.User) error {
	return config.DB.Create(user).Error
}

func GetUsers() ([]model.User, error) {

	var users []model.User

	result := config.DB.Find(&users)

	return users, result.Error
}

func GetUserByID(id string) (model.User, error) {

	var user model.User

	userID, err := strconv.Atoi(id)
	if err != nil {
		return user, err
	}

	result := config.DB.First(&user, userID)

	return user, result.Error
}

func UpdateUser(user *model.User) error {

	result := config.DB.Save(user)

	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}

	return result.Error
}

func DeleteUser(id string) error {

	userID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	result := config.DB.
		Unscoped().
		Delete(&model.User{}, userID)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}
