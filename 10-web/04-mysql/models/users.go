package models

import (
	"fmt"
	"webmysql/db"
)

type User struct {
	ID       int64
	Username string
	Password string
	Email    string
}

type Users []User

const UserSchema string = `CREATE TABLE users (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

func NewUser(name, password, email string) *User {
	user := &User{
		Username: name,
		Password: password,
		Email:    email,
	}
	return user
}

// crear usuario e insertar
func CreateUser(name, password, email string) *User {
	user := NewUser(name, password, email)
	user.insert()
	return user
}

// insertar registro
func (u *User) insert() {
	sql := "INSERT INTO users SET username=?, password=?, email=?"
	r, err := db.Exec(sql, u.Username, u.Password, u.Email)
	if err != nil {
		fmt.Println("ERROR: Insert into, ", err.Error())
	}
	// fmt.Println(r)
	u.ID, _ = r.LastInsertId()
}

// listar todos los registros
func ListUsers() Users {
	sql := "SELECT id, username, password, email FROM users"
	users := Users{}
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("ERROR: list users, ", err.Error())
	}

	for rows.Next() {
		user := User{}
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}

	return users
}

// get a user
func GetUser(id int64) *User {
	user := NewUser("", "", "")
	sql := "SELECT id, username, password, email FROM users WHERE id=?"
	rows, _ := db.Query(sql, id)

	for rows.Next() {
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	}
	return user
}

func (u *User) update() {
	sql := "Update users SET username=?, password=?, email=? WHERE id=?"
	_, err := db.Exec(sql, u.Username, u.Password, u.Email, u.ID)
	if err != nil {
		fmt.Println("ERROR: update user, ", err.Error())
	}
}

func (u *User) Save() {
	if u.ID == 0 {
		u.insert()
	} else {
		u.update()
	}
}

func (u *User) Delete() {
	sql := "DELETE FROM users WHERE id=?"
	_, err := db.Exec(sql, u.ID)
	if err != nil {
		fmt.Println("ERROR: delete user, ", err.Error())
	}
}
