package app

import "github.com/inkochetkov/auth/internal/entity"

type API interface {
	// get user
	GetEntity(conditional map[string]any) (*entity.User, error)
	// get list user
	ListEntity() ([]*entity.User, error)
	// user operations
	ChangeEntity(user *entity.User, operation string) error
}
