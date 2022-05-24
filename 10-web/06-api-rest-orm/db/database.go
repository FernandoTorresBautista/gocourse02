package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:12345@tcp(localhost:3306)/goweb?charset=utf8mb4&parseTime=True&loc=Local"
var Database = func() (db *gorm.DB) {
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("ERROR: connection to db, ", err.Error())
		panic(err)
	} else {
		fmt.Println("Conexión exitosa")
		return db
	}
}()
