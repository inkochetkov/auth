package sqlite

import (
	"github.com/inkochetkov/auth/internal/entity"
)

// Create entity
func (b *SQLite) Create(user *entity.User) error {

	err := b.db.Create(user).Error
	if err != nil {
		b.logger.Error("Create", err)
		return err
	}

	return nil
}
