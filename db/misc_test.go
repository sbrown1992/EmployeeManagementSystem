package db

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func establishDBConnection(t *testing.T) *DB {
	var server = "0.0.0.0"
	var port = 1433
	var user = "sa"
	var database = "Master"

	godotenv.Load("../.env")
	password := os.Getenv("MSSQL_SA_PASSWORD")

	db, err := Connect(server, port, user, password, database)
	require.Nil(t, err)

	return db
}
