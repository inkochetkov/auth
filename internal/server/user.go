package server

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/inkochetkov/auth/internal/entity"
	"github.com/inkochetkov/gen-str"
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

	userDB := &entity.User{
		Login:    *user.Login,
		Password: password,
	}

	if user.Option != nil {
		userDB.Option, err = entity.SetOption(*user.Option)
		if err != nil {
			renderError(c, http.StatusMethodNotAllowed, err)
			return
		}
	}

	err = r.Access(c, entity.Create, entity.Zero)
	if err != nil {
		renderError(c, http.StatusMethodNotAllowed, err)
		return
	}

	err = r.api.ChangeEntity(userDB, entity.Create)
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

	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		renderError(c, http.StatusMethodNotAllowed, err)
		return
	}

	userDB := &entity.User{ID: id}

	if user.Login != nil {
		userDB.Login = *user.Login
	}

	if user.Password != nil {
		password, err := gen.GenPassword(*user.Password)
		if err != nil {
			renderError(c, http.StatusMethodNotAllowed, err)
			return
		}
		userDB.Password = password
	}

	if user.Option != nil {
		userDB.Option, err = entity.SetOption(*user.Option)
		if err != nil {
			renderError(c, http.StatusMethodNotAllowed, err)
			return
		}
	}

	err = r.Access(c, entity.Update, id)
	if err != nil {
		renderError(c, http.StatusMethodNotAllowed, err)
		return
	}

	err = r.api.ChangeEntity(userDB, entity.Update)
	if err != nil {
		renderError(c, http.StatusMethodNotAllowed, err)
		return
	}
}

// Delete entity user
func (r *Router) Delete(c *gin.Context, ids string) {

	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		renderError(c, http.StatusMethodNotAllowed, err)
		return
	}

	err = r.Access(c, entity.Delete, id)
	if err != nil {
		renderError(c, http.StatusMethodNotAllowed, err)
		return
	}

	err = r.api.ChangeEntity(&entity.User{ID: id}, entity.Delete)
	if err != nil {
		renderError(c, http.StatusMethodNotAllowed, err)
		return
	}
}

// Get entity user
func (r *Router) Get(c *gin.Context, ids string) {

	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		renderError(c, http.StatusMethodNotAllowed, err)
		return
	}

	err = r.Access(c, entity.Get, id)
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

	err := r.Access(c, entity.GetList, entity.Zero)
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
