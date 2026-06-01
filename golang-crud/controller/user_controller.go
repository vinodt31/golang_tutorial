package controller

import (
	"encoding/json"
	"fmt"
	"golang-crud/model"
	"golang-crud/service"
	"net/http"
	"strings"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": err.Error(),
			},
		)
		return
	}

	err = service.CreateUser(&user)

	if err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": err.Error(),
			},
		)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		user,
	)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := service.GetUsers()

	if err != nil {
		writeJSON(
			w,
			http.StatusInternalServerError,
			map[string]string{
				"error": err.Error(),
			},
		)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		users,
	)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/users/")

	user, err := service.GetUserByID(id)

	if err != nil {
		writeJSON(
			w,
			http.StatusNotFound,
			map[string]string{
				"error": "User not found",
			},
		)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		user,
	)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/users/")

	user, err := service.GetUserByID(id)

	if err != nil {
		writeJSON(
			w,
			http.StatusNotFound,
			map[string]string{
				"error": "User not found",
			},
		)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": err.Error(),
			},
		)
		return
	}

	err = service.UpdateUser(&user)

	if err != nil {
		writeJSON(
			w,
			http.StatusInternalServerError,
			map[string]string{
				"error": err.Error(),
			},
		)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		user,
	)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/users/")
	fmt.Println("id : ", id)

	err := service.DeleteUser(id)

	if err != nil {
		writeJSON(
			w,
			http.StatusInternalServerError,
			map[string]string{
				"error": err.Error(),
			},
		)
		return
	}

	response := map[string]string{
		"message": "User deleted successfully",
	}

	writeJSON(
		w,
		http.StatusOK,
		response,
	)
}
