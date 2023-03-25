package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inkochetkov/auth/internal/entity"
)

func NewHTTP(cfg entity.Config, router *Router) http.Handler {

	gin.SetMode(cfg.HTTP.Mode)

	r := gin.New()

	r.Any("/", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	r.Use(router.CheckAuth)

	return RegisterHandlers(r, router)
}
