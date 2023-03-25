package sqlite

import (
	"log"

	"github.com/inkochetkov/auth/internal/entity"
	"github.com/jmoiron/sqlx"
)

// SQLite ...
type SQLite struct {
	cfg  entity.Config
	conn *sqlx.DB
}

const (
	driverNameSqLite3 = "sqlite3"
)

// New SQLite
func New(cfg entity.Config) *SQLite {

	url, err := checkFileBD(cfg)
	if err != nil {
		log.Fatal(driverNameSqLite3, err)
		return nil
	}

	dbConnection, err := sqlx.Open(driverNameSqLite3, url)
	if err != nil {
		log.Fatal(driverNameSqLite3, err)
		return nil
	}

	return &SQLite{
		cfg:  cfg,
		conn: dbConnection,
	}
}
