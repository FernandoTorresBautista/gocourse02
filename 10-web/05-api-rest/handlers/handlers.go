package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"webapirest/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// // w.Header().Set("Content-Type", "text/xml")
	// // yaml no tiene un content type especifico
	// db.Connect()
	// users, _ := models.ListUsers()
	// db.Close()
	// bdata, err := json.Marshal(users)
	// // bdata, err := xml.Marshal(users)
	// // bdata, err := yaml.Marshal(users)
	// if err != nil {
	// 	fmt.Println("ERROR: getusers, ", err.Error())
	// }
	// fmt.Fprintln(w, string(bdata))
	if users, err := models.ListUsers(); err != nil {
		models.SendNotFound(w, nil)
	} else {
		models.SendData(w, users)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// // Obteer ID
	// // // way 1
	// // // fmt.Println(r.URL.Path)
	// // p := strings.Split(r.URL.Path, "/")
	// // // fmt.Println(p, len(p))
	// // if len(p) != 4 {
	// // 	fmt.Fprintln(w, "Bad request")
	// // }
	// // idx, _ := strconv.Atoi(p[3])
	// // fmt.Println("ID: ", idx)
	// // way 2, using mux
	// vars := mux.Vars(r)
	// idx, _ := strconv.Atoi(vars["id"])

	// db.Connect()
	// user, _ := models.GetUser(int64(idx))
	// db.Close()

	// bdata, err := json.Marshal(user)
	// if err != nil {
	// 	fmt.Println("ERROR: getuser, ", err.Error())
	// }
	// fmt.Fprintln(w, string(bdata))
	if user, err := getUserByRequest(r); err != nil {
		models.SendNotFound(w, nil)
	} else {
		models.SendData(w, user)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// obtener regustro
	user := models.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(w, nil)
		return
	}
	// db.Connect()
	// user.Save()
	// db.Close()
	// bdata, err := json.Marshal(user)
	// if err != nil {
	// 	fmt.Println("ERROR: CreateUser, ", err.Error())
	// }
	// fmt.Fprintln(w, string(bdata))
	user.Save()
	models.SendData(w, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	var userId int64
	if user, err := getUserByRequest(r); err != nil {
		models.SendNotFound(w, nil)
	} else {
		userId = user.ID
	}

	// obtener regustro
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(w, nil)
		return
	}

	// db.Connect()
	// user.Save()
	// db.Close()
	// bdata, err := json.Marshal(user)
	// if err != nil {
	// 	fmt.Println("ERROR: CreateUser, ", err.Error())
	// }
	// fmt.Fprintln(w, string(bdata))
	user.ID = userId
	user.Save()
	models.SendData(w, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	// vars := mux.Vars(r)
	// idx, _ := strconv.Atoi(vars["id"])

	// db.Connect()
	// user, _ := models.GetUser(int64(idx))
	// user.Delete()
	// db.Close()

	// bdata, err := json.Marshal(user)
	// if err != nil {
	// 	fmt.Println("ERROR: getuser, ", err.Error())
	// }
	// fmt.Fprintln(w, string(bdata))
	if user, err := getUserByRequest(r); err != nil {
		models.SendNotFound(w, nil)
	} else {
		user.Delete()
		models.SendData(w, user)
	}
}

func getUserByRequest(r *http.Request) (models.User, error) {
	vars := mux.Vars(r)
	idx, _ := strconv.Atoi(vars["id"])

	if user, err := models.GetUser(int64(idx)); err != nil {
		return *user, err
	} else {
		return *user, nil
	}
}
