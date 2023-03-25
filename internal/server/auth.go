package server

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/inkochetkov/auth/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

func (r *Router) Login(c *gin.Context) {

	userClient := LoginData{}
	if err := c.BindJSON(&userClient); err != nil {
		renderError(c, http.StatusUnauthorized, err)
		return
	}

	conditional := map[string]any{
		"login = ? ": *userClient.Login,
	}

	user, err := r.api.GetEntity(conditional)
	if err != nil {
		renderError(c, http.StatusUnauthorized, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(*userClient.Password))
	if err != nil {
		renderError(c, http.StatusUnauthorized, err)
		return
	}

	tokenExpiresTime := jwt.NumericDate{Time: time.Now().UTC().Add(r.cfg.JWT.TTL)}

	tokenClaims := &entity.AuthClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &tokenExpiresTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	t, err := token.SignedString([]byte(r.cfg.JWT.TokenClaims))
	if err != nil {
		renderError(c, http.StatusUnauthorized, err)
		return
	}

	items := map[string]any{"token": t}
	err = r.api.EntityChange(items, conditional, entity.Update)
	if err != nil {
		renderError(c, http.StatusUnauthorized, err)
		return
	}

	renderResponse(c, http.StatusOK, map[string]string{"Token": t})
}
func (r *Router) Logout(c *gin.Context) {

	token, ok := c.Get(entity.Token)
	if !ok {
		renderError(c, http.StatusUnauthorized, errors.New("token user fail"))
		return
	}

	conditional := map[string]any{
		entity.Token: token,
	}

	err := r.api.EntityChange(nil, conditional, entity.Update)
	if err != nil {
		renderError(c, http.StatusUnauthorized, err)
		return
	}

	renderResponse(c, http.StatusOK, nil)

}
