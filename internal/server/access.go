package server

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/inkochetkov/auth/internal/entity"
)

func (r *Router) Access(c *gin.Context, operationID string, id int) error {

	token, ok := c.Get(entity.Token)
	if !ok {
		if operationID == entity.Create {
			return nil
		}
		return errors.New("token user fail")
	}

	user, err := r.api.GetEntity(map[string]any{
		"token = ?": token,
	})
	if err != nil {
		return errors.New("user fail")
	}

	if operationID != entity.GetList &&
		user.ID == id {
		return nil
	}

	role, ok := user.Option["role"]
	if !ok {
		return errors.New("role is empty")
	}

	if role != "admin" {
		return errors.New("didn`t access")
	}

	return nil

}
