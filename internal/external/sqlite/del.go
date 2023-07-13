package sqlite

import (
	"github.com/inkochetkov/auth/internal/entity"
)

// Delete entity
func (b *SQLite) Delete(user *entity.User) error {

	err := b.db.Where("id = ?", user.ID).Delete(&entity.User{}).Error
	if err != nil {
		b.logger.Error("Delete", err)
		return err
	}
	return nil
}
