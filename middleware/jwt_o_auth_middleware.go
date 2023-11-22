package middleware

import (
	"errors"
	"net/http"
	"os"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/helpers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ChekLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		// Get jwt token from cookie
		tokenStr, err := c.Cookie("token")

		// check if cookie is exist
		if err != nil {
			// if not clear session and send response api
			session.Clear()
			session.Save()
			err = errors.New("Unauthorized!")
			response := helpers.UserResponseFail(err)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		claims := &app.JwtClaims{}

		// parse token in cookie
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
			key := os.Getenv("JWT_KEY")
			return []byte(key), nil
		})
		// if cookie value does not match
		// send response api unauthorize
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				session.Clear()
				session.Save()
				err = errors.New("Unauthorized!")
				response := helpers.UserResponseFail(err)
				c.JSON(http.StatusUnauthorized, response)
				c.Abort()
				return
			}
			response := helpers.UserResponseFail(err)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}
		// if token is expired send response api un authorize
		if !token.Valid {
			session.Clear()
			session.Save()
			response := helpers.UserResponseFail(err)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}
		c.Next()
	}

}
