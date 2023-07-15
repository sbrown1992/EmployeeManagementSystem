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

	// TODO move these elsewhere
	err = db.CleanDB(conn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.SetupDB(conn)
	if err != nil {
		log.Fatal(err)
	}
}
