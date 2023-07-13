package sqlite

import (
	"errors"

	"github.com/inkochetkov/auth/internal/entity"
)

// Update entity
func (b *SQLite) Update(user *entity.User) error {

	result := b.db.Model(&entity.User{}).Where("id =?", user.ID).Updates(user)
	if result.RowsAffected == 0 {
		return errors.New("user data not update")
	}

	return nil
}
