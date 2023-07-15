package db

import (
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/microsoft/go-mssqldb"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type DB struct {
	conn        *gorm.DB
	scriptsPath string
}

func Connect(server string, port int, user, password, database string) (*DB, error) {
	connString := fmt.Sprintf("server=%s;user id = %s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	conn, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db := &DB{
		conn:        conn,
		scriptsPath: "../db_scripts",
	}
	return db, nil
}

func (db *DB) ResetDB() error {
	err := db.CleanDB()
	if err != nil {
		return err
	}
	return db.SetupDB()
}

func (db *DB) CleanDB() error {
	return db.runSQLScript(filepath.Join(db.scriptsPath, "clean.sql"))
}

func (db *DB) SetupDB() error {
	return db.runSQLScript(filepath.Join(db.scriptsPath, "setup.sql"))
}

func (db *DB) runSQLScript(filename string) error {
	queryBytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	query := string(queryBytes)
	db.conn.Exec(query)
	return nil
}

func (db *DB) SetScriptsPath(path string) {
	db.scriptsPath = path
}
