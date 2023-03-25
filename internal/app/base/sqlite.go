package base

import (
	"context"

	"github.com/inkochetkov/auth/internal/entity"
)

// GetEntity  user
func (a *API) GetEntity(conditional map[string]any) (*entity.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), a.config.SQL.Timeout)
	defer cancel()

	var conditions string
	var values []any

	for name, value := range conditional {
		conditions = name
		values = append(values, value)
	}

	userDB, err := a.sql.Get(ctx, conditions, values)
	if err != nil {
		return nil, err
	}

	option, err := unmarshalJSONToMap(userDB.Option)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		ID:       userDB.ID,
		Login:    userDB.Login,
		Password: userDB.Password,
		Option:   option,
	}

	if userDB.Token.Valid {
		user.Token = userDB.Token.String
	}

	return user, nil
}

// ListEntity  users
func (a *API) ListEntity() ([]*entity.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), a.config.SQL.Timeout)
	defer cancel()

	usersDB, err := a.sql.GetList(ctx)
	if err != nil {
		return nil, err
	}

	var users []*entity.User

	for _, userDB := range usersDB {

		option, err := unmarshalJSONToMap(userDB.Option)
		if err != nil {
			return nil, err
		}

		user := &entity.User{
			ID:     userDB.ID,
			Login:  userDB.Login,
			Option: option,
		}
		if userDB.Token.Valid {
			user.Token = userDB.Token.String
		}

		users = append(users, user)
	}

	return users, nil
}

// ChangeEntity - entity operations
func (a *API) ChangeEntity(items, condition map[string]any, operation string) error {

	ctx, cancel := context.WithTimeout(context.Background(), a.config.SQL.Timeout)
	defer cancel()

	switch operation {
	case entity.Create:
		err := a.sql.Create(ctx, items)
		if err != nil {
			return err
		}
	case entity.Delete:
		err := a.sql.Delete(ctx, condition)
		if err != nil {
			return err
		}
	case entity.Update:
		err := a.sql.Update(ctx, items, condition)
		if err != nil {
			return err
		}
	}

	return nil
}
