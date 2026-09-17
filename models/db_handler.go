package models

import (
	"os"
	"database/sql"
	"path/filepath"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

func NewDB() (*DB, error) {
	db, err := sql.Open("sqlite3", getDBPath())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return &DB{db}, nil
}

func getDBPath() string {
	if envPath := os.Getenv("TSP_DB_PATH"); envPath != "" {
		return envPath
	}
	
	execPath, err := os.Executable()
	if err != nil {
		return filepath.Join("data", "tsp.db")
	}

	execDir := filepath.Dir(execPath)

	return filepath.Join(execDir, "data", "tsp.db")
}
