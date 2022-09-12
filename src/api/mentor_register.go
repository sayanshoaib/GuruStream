package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"softwareLab/src/model"
)

func MentorRegister(c echo.Context) error {
	return c.Render(http.StatusOK, "mentorForm", nil)
}

func MentorRegisterProcess(c echo.Context) error {
	name := c.FormValue("name")
	worksAt := c.FormValue("work")
	expertise := c.FormValue("expertise")
	designation := c.FormValue("designation")
	phone := c.FormValue("phone")
	email := c.FormValue("email")
	password := c.FormValue("pass")

	file, err := c.FormFile("image")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	encryptedPassword, err := HashPassword(password)

	mentor := model.Mentor{
		Name:            name,
		ProfilePicture:  file.Filename,
		Email:           email,
		WorksAt:         worksAt,
		Phone:           phone,
		Designation:     designation,
		FieldOfInterest: expertise,
		Password:        encryptedPassword,
	}
	fmt.Println(mentor)
	result := model.DB.Create(&mentor)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(result)

	return c.Redirect(http.StatusTemporaryRedirect, "/mentor/dashboard")
}
