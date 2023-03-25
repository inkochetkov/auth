package server

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/inkochetkov/auth/internal/entity"
)

const (
	authHeader         = "Authorization"
	bearerPrefix       = "Bearer "
	bearerPrefixLength = len(bearerPrefix)
)

func (r *Router) CheckAuth(c *gin.Context) {

	if c.Request.RequestURI == "/login/" ||
		c.Request.RequestURI == "/logout/" ||
		c.Request.RequestURI == "/check/" {
		return
	}

	header := c.GetHeader(authHeader)

	tokenString, err := extractJWT(header)
	if err != nil {
		renderError(c, http.StatusUnauthorized, err)
		return
	}

	tokenClaims, err := jwt.ParseWithClaims(tokenString, &entity.AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(r.cfg.JWT.TokenClaims), nil
	})
	if err != nil {
		renderError(c, http.StatusUnauthorized, err)
		return
	}

	if !tokenClaims.Valid {
		renderError(c, http.StatusUnauthorized, errors.New("not valid jwt"))
		return
	}

	c.Set(entity.Token, tokenString)

}

func extractJWT(header string) (string, error) {

	var possibleJWTString string

	if strings.HasPrefix(header, bearerPrefix) {
		possibleJWTString = strings.TrimPrefix(header, bearerPrefix)
	} else {
		if !strings.HasPrefix(header, bearerPrefix) {
			return entity.Empty, errors.New("no bearer in authorization header")
		}
		possibleJWTString = header[bearerPrefixLength:]
	}

	jwtString := strings.TrimSpace(possibleJWTString)
	if jwtString == entity.Empty {
		return entity.Empty, errors.New("missing jwt")
	}
	return jwtString, nil
}
