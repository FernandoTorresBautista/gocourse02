package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// username:password@tcp(localhost:3306)/database

const url = "root:12345@tcp(localhost:3306)/goweb"

// save connection
var db *sql.DB

// connect
func Connect() {
	conn, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexión exitosa")
	db = conn
}

// close connection
func Close() {
	db.Close()
}

// check the connection
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

// check if some table exists
func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("ERROR: Exists table, ", err.Error())
	}
	return rows.Next()
}

// create table(s)
func CreateTable(schema, name string) {
	if ExistsTable(name) {
		return
	}
	r, err := db.Exec(schema)
	if err != nil {
		fmt.Println("ERROR: Creating table, " + err.Error())
	}
	fmt.Println("table created", r)
}

// polimorfismo de Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	result, err := db.Exec(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

// polimorfismo de Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect()
	rows, err := db.Query(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}

// reiniciar todos los registros de una tabla
func TruncateTable(name string) {
	sql := fmt.Sprintf("TRUNCATE %s", name)
	Exec(sql)
}
