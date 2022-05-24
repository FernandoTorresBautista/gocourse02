package models

import "webapirestgorm/db"

type User struct {
	ID       int64  `json:"id" xml:"id" yaml:"id"`
	Username string `json:"username" xml:"username" yaml:"username"`
	Password string `json:"password" xml:"password" yaml:"password"`
	Email    string `json:"email" xml:"email" yaml:"email"`
}

type Users []User

func MigrateUser() {
	db.Database.AutoMigrate(User{})
}
