package sqlite

import (
	"database/sql"
	"sync"

	p "path"

	"github.com/inkochetkov/auth/internal/entity"
	"github.com/inkochetkov/exist"
	"github.com/inkochetkov/gen-str"
	"github.com/inkochetkov/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SQLite ...
type SQLite struct {
	mu     sync.Mutex
	db     *gorm.DB
	conn   *sql.DB
	logger *log.Log
}

const (
	driverNameSqLite3 = "sqlite3"
)

// New SQLite
func New(conf entity.Config, logger *log.Log) *SQLite {

	dataSourcePath, err := checkFileBD(conf.SQL.Dir, conf.SQL.Name)
	if err != nil {
		logger.Fatal(entity.SqliteNameModule, "path to bd fail", err)
	}

	dbGorm, err := gorm.Open(sqlite.Open(dataSourcePath), &gorm.Config{})
	if err != nil {
		logger.Fatal(entity.SqliteNameModule, "connect fail", err)
	}

	conn, err := dbGorm.DB()
	if err != nil {
		logger.Fatal(entity.SqliteNameModule, "connect fail", err)
	}

	err = dbGorm.AutoMigrate(&entity.User{})
	if err != nil {
		logger.Fatal(entity.SqliteNameModule, "AutoMigrate fail", err)
	}

	checkUserFirst(conf, dbGorm, logger)

	return &SQLite{conn: conn, db: dbGorm, logger: logger}
}

// checkFileBD - check dir and file exist
func checkFileBD(path, fileName string) (string, error) {

	url := p.Join(path, fileName)

	if ok := exist.CheckFile(url); ok {
		return url, nil
	}

	_, err := exist.InitDirFile(path, fileName)
	if err != nil {
		return entity.Empty, err
	}

	return url, nil
}

// check find first user
func checkUserFirst(conf entity.Config, db *gorm.DB, logger *log.Log) {

	var users []*entity.User
	err := db.Find(&users).Error
	if err != nil {
		logger.Fatal("checkUserFirst", err)
	}

	if len(users) != 0 {
		return
	}

	password, err := gen.GenPassword(conf.FirstUser.Password)
	if err != nil {
		logger.Fatal("checkUserFirst,GenPassword", err)
	}
	option, err := entity.SetOption(map[string]any{"role": "admin"})
	if err != nil {
		logger.Fatal("checkUserFirst, SetOption", err)
	}
	user := &entity.User{
		Login:    conf.FirstUser.Login,
		Password: password,
		Option:   option,
	}

	err = db.Create(user).Error
	if err != nil {
		logger.Fatal("checkUserFirst", err)
	}

}
