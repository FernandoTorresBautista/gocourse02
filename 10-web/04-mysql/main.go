package main

import (
	"fmt"

	"webmysql/db"
	"webmysql/models"
)

func main() {
	db.Connect()

	db.Ping()

	db.CreateTable(models.UserSchema, "users")
	// db.TruncateTable("users")

	// user := models.CreateUser("Fernando2", "12345", "test2@test.com")
	// fmt.Println(user)

	// users := models.ListUsers()
	// fmt.Println(users)

	// user := models.GetUser(1)
	// user.Username = "Fernando modificado 2"
	// user.Save()
	// fmt.Println(user)

	// user := models.GetUser(2)
	// fmt.Println(user)
	// user.Delete()

	fmt.Println(models.ListUsers())

	db.Close()
}
