package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"webapirestgorm/db"
	"webapirestgorm/models"

	"gorm.io/gorm"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := models.Users{}
	db.Database.Find(&users)
	sendData(w, users, http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := getUserByID(r)
	if err != nil {
		sendError(w, http.StatusNotFound)
		return
	}
	sendData(w, user, http.StatusOK)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		sendError(w, http.StatusUnprocessableEntity)
		return
	}
	db.Database.Save(&user)
	sendData(w, user, http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var userId int64
	user_ant, err := getUserByID(r)
	if err != nil {
		sendError(w, http.StatusNotFound)
		return
	}
	userId = user_ant.ID

	// obtener regustro
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		sendError(w, http.StatusUnprocessableEntity)
		return
	}
	user.ID = userId
	db.Database.Save(&user)
	sendData(w, user, http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	user, err := getUserByID(r)
	if err != nil {
		sendError(w, http.StatusNotFound)
		return
	}
	db.Database.Delete(&user)
	sendData(w, user, http.StatusOK)
}

func getUserByID(r *http.Request) (models.User, *gorm.DB) {
	vars := mux.Vars(r)
	idx, _ := strconv.Atoi(vars["id"])

	user := models.User{}
	if err := db.Database.First(&user, idx); err.Error != nil {
		return user, err
	}

	return user, nil
}
