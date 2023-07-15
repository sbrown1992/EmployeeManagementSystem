package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/microsoft/go-mssqldb"
)

type DB struct {
	conn *sql.DB
}

func Connect(server string, port int, user, password, database string) (*DB, error) {
	connString := fmt.Sprintf("server=%s;user id = %s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	conn, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()

	err = conn.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	db := &DB{
		conn: conn,
	}
	return db, nil
}

func (db *DB) CleanDB() error {
	return db.runSQLScript("db_scripts/clean.sql")
}

func (db *DB) SetupDB() error {
	return db.runSQLScript("db_scripts/setup.sql")
}

func (db *DB) runSQLScript(filename string) error {
	queryBytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	query := string(queryBytes)
	_, err = db.conn.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
