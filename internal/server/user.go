package server

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/inkochetkov/auth/internal/entity"
	"github.com/inkochetkov/gen-str/pkg/gen"
)

// Create entity user
func (r *Router) Create(c *gin.Context) {

	var user User
	err := c.BindJSON(&user)
	if err != nil {
		renderError(c, http.StatusMethodNotAllowed, err)
		return
	}

	if user.Login == nil || user.Password == nil {
		renderError(c, http.StatusMethodNotAllowed, errors.New("didn`t credintional"))
		return
	}

	password, err := gen.GenPassword(*user.Password)
	if err != nil {
		renderError(c, http.StatusMethodNotAllowed, err)
		return
	}

	items := map[string]any{
		"login":    user.Login,
		"password": password,
	}

	if user.Option != nil {
		items["option"] = *user.Option
	}

	err = r.Access(c, entity.Create)
	if err != nil {
		renderError(c, http.StatusMethodNotAllowed, err)
		return
	}

	err = r.api.EntityChange(items, nil, entity.Create)
	if err != nil {
		renderError(c, http.StatusMethodNotAllowed, err)
		return
	}
}

// Update entity user
func (r *Router) Update(c *gin.Context, ids string) {

	var user User
	err := c.BindJSON(&user)
	if err != nil {
		renderError(c, http.StatusMethodNotAllowed, err)
		return
	}

	id, err := strconv.Atoi(ids)
	if err != nil {
		renderError(c, http.StatusMethodNotAllowed, err)
		return
	}

	items := make(map[string]any)

	if user.Login != nil {
		items["login"] = *user.Login
	}

	if user.Password != nil {
		password, err := gen.GenPassword(*user.Password)
		if err != nil {
			renderError(c, http.StatusMethodNotAllowed, err)
			return
		}
		items["password"] = password
	}

	if user.Option != nil {
		items["option"] = *user.Option
	}

	condition := map[string]any{
		"id = ?": id,
	}

	err = r.Access(c, entity.Update)
	if err != nil {
		renderError(c, http.StatusMethodNotAllowed, err)
		return
	}

	err = r.api.EntityChange(items, condition, entity.Update)
	if err != nil {
		renderError(c, http.StatusMethodNotAllowed, err)
		return
	}
}

// Delete entity user
func (r *Router) Delete(c *gin.Context, ids string) {

	id, err := strconv.Atoi(ids)
	if err != nil {
		renderError(c, http.StatusMethodNotAllowed, err)
		return
	}

	err = r.Access(c, entity.Delete)
	if err != nil {
		renderError(c, http.StatusMethodNotAllowed, err)
		return
	}

	condition := map[string]any{
		"id = ?": id,
	}

	err = r.api.EntityChange(nil, condition, entity.Delete)
	if err != nil {
		renderError(c, http.StatusMethodNotAllowed, err)
		return
	}
}

// Get entity user
func (r *Router) Get(c *gin.Context, ids string) {

	id, err := strconv.Atoi(ids)
	if err != nil {
		renderError(c, http.StatusMethodNotAllowed, err)
		return
	}

	err = r.Access(c, entity.Get)
	if err != nil {
		renderError(c, http.StatusMethodNotAllowed, err)
		return
	}

	condition := map[string]any{
		"id = ?": id,
	}

	user, err := r.api.GetEntity(condition)
	if err != nil {
		renderError(c, http.StatusMethodNotAllowed, err)
		return
	}

	user.Password = ""

	renderResponse(c, http.StatusOK, user)
}

// List entity users
func (r *Router) List(c *gin.Context) {
	err := r.Access(c, entity.GetList)
	if err != nil {
		renderError(c, http.StatusMethodNotAllowed, err)
		return
	}

	users, err := r.api.ListEntity()
	if err != nil {
		renderError(c, http.StatusMethodNotAllowed, err)
		return
	}

	renderResponse(c, http.StatusOK, users)
}
