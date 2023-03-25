package server

import "github.com/gin-gonic/gin"

func renderResponse(c *gin.Context, status int, entity any) {
	c.JSON(status, entity)
}

func renderError(c *gin.Context, status int, err error) {
	c.JSON(status, err.Error())
}
