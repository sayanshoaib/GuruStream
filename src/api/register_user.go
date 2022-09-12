package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"softwareLab/src/model"
)

func RegisterView(c echo.Context) error {
	message := c.QueryParam("message")
	return c.Render(http.StatusOK, "register", map[string]interface{}{
		"error": message,
	})
}

func RegisterUser(c echo.Context) error {
	uName := c.FormValue("username")
	uEmail := c.FormValue("email")
	uPassword := c.FormValue("password")
	encryptedPassword, err := HashPassword(uPassword)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusTemporaryRedirect, "/register")
	}

	user := model.User{
		Name:     uName,
		Email:    uEmail,
		Password: encryptedPassword,
	}
	fmt.Println(user)
	result := model.DB.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(result)
	return c.Redirect(http.StatusTemporaryRedirect, "/mentee")
}
