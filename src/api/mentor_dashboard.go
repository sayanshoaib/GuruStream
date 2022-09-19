package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"softwareLab/src/model"
)

func MentorDashboard(c echo.Context) error {
	mentorId, err := c.Cookie("mentorId")
	if err != nil {
		fmt.Println(err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "/mentor/register?message=mentor_cookie_userId_not_found")
	}

	fmt.Println(mentorId.Name)
	fmt.Println(mentorId.Value)

	mentor := &model.Mentor{}

	tx := model.DB.First(mentor, "id = ?", mentorId.Value)
	if tx.Error != nil {
		fmt.Println("User not found")
		return c.Redirect(http.StatusTemporaryRedirect, "/mentor/register?message=User_not_found")
	}

	fmt.Println(mentor)

	return c.Render(http.StatusOK, "mentorDashboard", map[string]interface{}{
		"mentor": mentor,
	})
}
