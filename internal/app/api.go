package app

import "github.com/inkochetkov/auth/internal/entity"

type API interface {
	// get entity
	GetEntity(conditional map[string]any) (*entity.User, error)
	// get list entity
	ListEntity() ([]*entity.User, error)
	// entity operations
	EntityChange(items, condition map[string]any, operation string) error
}
