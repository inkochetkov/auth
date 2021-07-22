package main

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnect(t *testing.T) {
	db, _ := sql.Open("sqlite3", "init/bd.sq3")
	defer db.Close()
}
