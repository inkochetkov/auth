package sqlite

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
)

// Create entity
func (b *SQLite) Create(ctx context.Context, items map[string]any) error {

	var (
		names  []string
		values []any
	)

	for name, value := range items {
		names = append(names, name)
		values = append(values, value)
	}

	q, arg, err := sq.
		Insert("user").
		Columns(names...).
		Values(values...).
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
