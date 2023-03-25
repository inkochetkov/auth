package entity

import (
	"context"
	"database/sql"
)

type UserDB struct {
	ID       int            `db:"id"`
	Login    string         `db:"login"`
	Password string         `db:"password"`
	Token    sql.NullString `db:"token"`
	Option   []byte         `db:"option"`
}

type ExteranlSQL interface {
	Get(ctx context.Context, conditional string, values []any) (*UserDB, error)
	GetList(ctx context.Context) ([]*UserDB, error)
	Create(ctx context.Context, items map[string]any) error
	Update(ctx context.Context, items, condition map[string]any) error
	Delete(ctx context.Context, condition map[string]any) error
}
