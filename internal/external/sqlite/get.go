package sqlite

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/inkochetkov/auth/internal/entity"
	_ "github.com/mattn/go-sqlite3"
)

func (b *SQLite) Get(ctx context.Context, conditional string, values []any) (*entity.UserDB, error) {

	q, arg, err := sq.
		Select("login, password, token, option").
		From("user").
		Where(conditional, values...).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}
	user := &entity.UserDB{}
	err = b.conn.QueryRowxContext(ctx, q, arg...).StructScan(user)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (b *SQLite) GetList(ctx context.Context) ([]*entity.UserDB, error) {

	q, arg, err := sq.
		Select("login, password, token, option").
		From("user").
		OrderBy("id").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := b.conn.QueryxContext(ctx, q, arg...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	var users []*entity.UserDB

	for rows.Next() {
		user := &entity.UserDB{}
		err := rows.StructScan(user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)

	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}
