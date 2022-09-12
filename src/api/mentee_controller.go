package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"softwareLab/src/model"
)

func MenteeDashboard(c echo.Context) error {
	userId, err := c.Cookie("userId")
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/register")
	}
	user := &model.User{}
	tx := model.DB.First(user, "id = ?", userId.Value)
	if tx.Error != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/register")
	}

	return c.Render(http.StatusOK, "mentee", map[string]interface{}{
		"user": user,
	})
}

func MenteeAccount(c echo.Context) error {
	return c.Render(http.StatusOK, "menteeAccount", nil)
}
