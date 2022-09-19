package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"softwareLab/src/model"
	"time"
)

func MentorLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "mentorLogin", nil)
}

func MentorLoginProcess(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	mentor := &model.Mentor{}

	tx := model.DB.First(mentor, "email = ?", email)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		if tx.Error == gorm.ErrRecordNotFound {
			return c.Redirect(http.StatusTemporaryRedirect, "/mentor/register?message=user not registered")
		}
		return c.Redirect(http.StatusTemporaryRedirect, "/mentor/register?message="+tx.Error.Error())
	}

	hashedPassword := mentor.Password
	ok := CheckPasswordHash(password, hashedPassword)

	if !ok {
		return c.Redirect(http.StatusTemporaryRedirect, "/mentor/register?message=email and password miss match")
	}

	cookie := new(http.Cookie)
	cookie.Name = "mentorId"
	cookie.Value = fmt.Sprintf("%d", mentor.ID)
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	fmt.Println(cookie.Name)

	return c.Redirect(http.StatusTemporaryRedirect, "/mentor/dashboard")
}
