package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inkochetkov/auth/internal/app"
	"github.com/inkochetkov/auth/internal/entity"
)

// NewRouter ...
func NewRouter(api app.API, cfg entity.Config) *Router {
	return &Router{
		api: api,
		cfg: cfg,
	}
}

// Router ...
type Router struct {
	api app.API
	cfg entity.Config
}

func (r *Router) Check(c *gin.Context) {
	renderResponse(c, http.StatusOK, nil)
}
