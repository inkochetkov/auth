package base

import (
	"github.com/inkochetkov/auth/internal/entity"
)

// GetEntity  user
func (a *API) GetEntity(conditional map[string]any) (*entity.User, error) {

	var conditions string
	var values []any

	for name, value := range conditional {
		conditions = name
		values = append(values, value)
	}

	user, err := a.sql.Get(conditions, values)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// ListEntity  users
func (a *API) ListEntity() ([]*entity.User, error) {

	users, err := a.sql.GetList()
	if err != nil {
		return nil, err
	}

	return users, nil
}

// ChangeEntity - entity operations
func (a *API) ChangeEntity(user *entity.User, operation string) error {

	switch operation {
	case entity.Create:
		err := a.sql.Create(user)
		if err != nil {
			return err
		}
	case entity.Delete:
		err := a.sql.Delete(user)
		if err != nil {
			return err
		}
	case entity.Update:
		err := a.sql.Update(user)
		if err != nil {
			return err
		}
	}

	return nil
}
