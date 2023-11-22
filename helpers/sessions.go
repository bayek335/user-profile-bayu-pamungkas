package helpers

import (
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetUserLogin(c *gin.Context, ID int, username, email string) {
	session := sessions.Default(c)

	year, month, day := time.Now().Date()
	y := strconv.Itoa(year)
	m := month.String()
	d := strconv.Itoa(day)

	h := strconv.Itoa(time.Now().Hour())
	min := strconv.Itoa(time.Now().Minute())
	s := strconv.Itoa(time.Now().Second())

	session.Set("id", ID)
	session.Set("username", username)
	session.Set("email", email)
	session.Set("login_date", y+", "+m+", "+d)
	session.Set("login_time", h+":"+min+":"+s)
	session.Save()
}
