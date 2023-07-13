package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inkochetkov/auth/internal/app"
	"github.com/inkochetkov/auth/internal/entity"
	"github.com/inkochetkov/log"
)

// NewRouter ...
func NewRouter(api app.API, cfg entity.Config, log *log.Log) *Router {
	return &Router{
		api: api,
		cfg: cfg,
		log: log,
	}
}

// Router ...
type Router struct {
	api app.API
	cfg entity.Config
	log *log.Log
}

func (r *Router) Check(c *gin.Context) {
	renderResponse(c, http.StatusOK, nil)
}
