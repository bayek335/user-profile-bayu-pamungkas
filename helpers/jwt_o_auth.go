package helpers

import (
	"os"
	"time"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtOAuth(c *gin.Context, email string) error {

	key := []byte(os.Getenv("JWT_KEY"))
	expTime := time.Now().Add(6 * time.Hour)
	claims := &app.JwtClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)

	if err != nil {
		// err := errors.New("Internal server error")
		return err
	}

	c.SetCookie("token", tokenString, 3600*6, "/", "localhost", true, false)

	return nil
}
