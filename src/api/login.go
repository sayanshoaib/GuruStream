package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"softwareLab/src/model"
	"time"
)

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	user := &model.User{}
	tx := model.DB.First(user, "email = ?", email)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		if tx.Error == gorm.ErrRecordNotFound {
			return c.Redirect(http.StatusTemporaryRedirect, "/register?message=user not registered")
		}
		return c.Redirect(http.StatusTemporaryRedirect, "/register?message="+tx.Error.Error())
	}

	hashedPassword := user.Password
	ok := CheckPasswordHash(password, hashedPassword)

	if !ok {
		return c.Redirect(http.StatusTemporaryRedirect, "/register?message=email and password miss match")
	}

	cookie := new(http.Cookie)
	cookie.Name = "userId"
	cookie.Value = fmt.Sprintf("%d", user.ID)
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	return c.Redirect(http.StatusTemporaryRedirect, "/mentee")
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
