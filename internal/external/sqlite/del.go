package sqlite

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
)

func (b *SQLite) Delete(ctx context.Context, condition map[string]any) error {

	var conditions string
	var values []any

	for name, value := range condition {
		conditions += ", " + name
		values = append(values, value)
	}

	q, arg, err := sq.
		Delete("user").
		Where(conditions, values...).
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
