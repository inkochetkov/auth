package base

import (
	"github.com/inkochetkov/auth/internal/entity"
	"github.com/inkochetkov/auth/internal/external/sqlite"
)

func NewAPI(
	config entity.Config,
	sql *sqlite.SQLite,

) *API {
	return &API{
		config: config,
		sql:    sql,
	}
}

type API struct {
	config entity.Config
	sql    entity.ExteranlSQL
}
