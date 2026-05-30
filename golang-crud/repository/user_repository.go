package repository

import (
	"golang-crud/config"
	"golang-crud/model"
)

func CreateUser(user *model.User) error{
	result := config.DB.Create(user)

	return result.Error
}

func GetUsers() ([]model.User, error){
	var users []model.User

	result := config.DB.Find(&users)

	return users, result.Error
}

func GetUserByID(id string) (model.User, error){
	var user model.User

	result := config.DB.First(&user, id)

	return user, result.Error
}

func UpdateUser(user *model.User) error{
	result := config.DB.Save(user)

	return result.Error
}

func DeleteUser(id string) error {
	result := config.DB.Delete(&model.User{}, id)

	return result.Error
}