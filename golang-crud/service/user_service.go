package service

import (
	"errors"
	"golang-crud/model"
	"golang-crud/repository"
)

func CreateUser (user *model.User) error {

	if user.Name == "" {
	return errors.New("name is required")
	}

	if user.Email == "" {
	return errors.New("email is required")
	}

	return repository.CreateUser(user)

}

func GetUsers() ([]model.User, error) {
	return repository.GetUsers()
}

func GetUserByID(id string) (model.User, error) {
	return repository.GetUserByID(id)
}

func UpdateUser(user *model.User) error {
	return repository.UpdateUser(user)
}

func DeleteUser(id string) error {
	return repository.DeleteUser(id)
}