package sqlite

import (
	"encoding/json"
	"path"

	"github.com/inkochetkov/auth/internal/entity"
	"github.com/inkochetkov/exist/pkg/exist"
	"github.com/inkochetkov/gen-str/pkg/gen"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func checkFileBD(cfg entity.Config) (string, error) {

	url := path.Join(cfg.SQL.PathSQL, cfg.SQL.PathSQLName)

	if ok := exist.CheckFile(url); ok {
		return url, nil
	}

	_, err := exist.InitDirFile(cfg.SQL.PathSQL, cfg.SQL.PathSQLName)
	if err != nil {
		return entity.Empty, err
	}

	err = initSchema(url, cfg)
	if err != nil {
		return entity.Empty, err
	}

	return url, nil
}

func initSchema(url string, cfg entity.Config) error {

	conn, err := sqlx.Open(driverNameSqLite3, url)
	if err != nil {
		return err
	}

	query := `CREATE TABLE user (
		id       INTEGER PRIMARY KEY AUTOINCREMENT
						 NOT NULL,
		login    TEXT    NOT NULL,
		password TEXT    NOT NULL,
		token    TEXT  ,
		option   BLOB ); 
		
		INSERT INTO user ('login','password', 'option')
		values(?,?,?)
		`

	option, err := json.Marshal(map[string]any{"role": "admin"})
	if err != nil {
		return err
	}

	password, err := gen.GenPassword(cfg.FirstUser.Password)
	if err != nil {
		return err
	}

	arg := []any{cfg.FirstUser.Login, password, option}
	_, err = conn.Exec(query, arg...)
	if err != nil {
		return err
	}

	conn.Close()

	return nil
}
