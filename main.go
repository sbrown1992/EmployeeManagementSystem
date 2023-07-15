package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sbrown1992/EmployeeManagementSystem/db"
)

var server = "0.0.0.0"
var port = 1433
var user = "sa"
var database = "Master"

func main() {
	fmt.Println("tech challenge")
	godotenv.Load()
	password := os.Getenv("MSSQL_SA_PASSWORD")

	conn, err := db.Connect(server, port, user, password, database)
	if err != nil {
		log.Fatal(err)
	}
	conn.SetScriptsPath("./db_scripts")

	err = conn.CleanDB()
	if err != nil {
		log.Fatal(err)
	}

	err = conn.SetupDB()
	if err != nil {
		log.Fatal(err)
	}
}
