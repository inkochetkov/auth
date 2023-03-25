package sqlite

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
)

func (b *SQLite) Update(ctx context.Context, items, condition map[string]any) error {

	builder := sq.Update("user")

	var arg []any

	for name, value := range items {
		arg = append(arg, value)
		builder = builder.Set(name, nil)
	}

	for name, value := range condition {
		arg = append(arg, value)
		builder = builder.Where(name)
	}

	q, _, err := builder.
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	_, err = b.conn.ExecContext(ctx, q, arg...)
	if err != nil {
		return err
	}

	return nil
}
