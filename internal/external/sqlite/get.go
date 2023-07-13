package sqlite

import (
	"github.com/inkochetkov/auth/internal/entity"
	_ "github.com/mattn/go-sqlite3"
)

// Get entity
func (b *SQLite) Get(conditional string, values []any) (*entity.User, error) {

	var user *entity.User
	err := b.db.Where(conditional, values...).Find(&user).Error
	if err != nil {
		b.logger.Error("Get", err)
		return nil, err
	}

	return user, nil
}

// GetList entity
func (b *SQLite) GetList() ([]*entity.User, error) {

	var users []*entity.User
	err := b.db.Find(users).Error
	if err != nil {
		b.logger.Error("GetList", err)
		return nil, err
	}

	return users, nil
}
