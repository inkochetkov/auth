package server

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/inkochetkov/auth/internal/entity"
)

func (r *Router) Access(c *gin.Context, operationID string) error {

	token, ok := c.Get(entity.Token)
	if !ok {
		return errors.New("token user fail")
	}

	user, err := r.api.GetEntity(map[string]any{
		"token = ?": token,
	})

	if err != nil {
		return errors.New("user fail")
	}

	role, ok := user.Option["role"]
	if !ok {
		return errors.New("rbac is fail")
	}

	if role != "admin" {
		return errors.New("didn`t access")
	}

	return nil

}
