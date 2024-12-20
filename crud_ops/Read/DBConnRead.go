package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// connect to sqlite db
// hardcode the port, doesnt really matter if its in and .env runs locally anyway
func connectReadDB() (*sql.DB, error) {
	kubernetesConnection := ""
	db, err := sql.Open("sqlite3", kubernetesConnection)

	if err != nil {
		return nil, fmt.Errorf("ERROR IN CONNECTREADDB FUNCTION:\n %w", err)
	}

	return db, nil
}
