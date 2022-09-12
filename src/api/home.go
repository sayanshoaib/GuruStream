package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Home(c echo.Context) error {
	//userId, err := c.Cookie("userId")
	//if err != nil {
	//	return c.Redirect(http.StatusTemporaryRedirect, "/register")
	//}
	//user := &model.User{}
	//tx := model.DB.First(user, "id = ?", userId.Value)
	//if tx.Error != nil {
	//	return c.Redirect(http.StatusTemporaryRedirect, "/register")
	//}
	//
	//return c.Render(http.StatusOK, "home", map[string]interface{}{
	//	"user": user,
	//})
	return c.Render(http.StatusOK, "home", nil)
}
