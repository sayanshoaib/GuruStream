package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"softwareLab/src/model"
)

func MenteeAccountProcess(c echo.Context) error {

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("profilePicture")
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

	language := c.FormValue("language")
	worksAt := c.FormValue("work")
	phone := c.FormValue("phone")
	designation := c.FormValue("designation")
	dob := c.FormValue("dob")
	location := c.FormValue("location")

	var fieldOfInterest string
	if c.FormValue("development") == "Development" {
		fieldOfInterest = "Development"
	} else if c.FormValue("design") == "Design" {
		fieldOfInterest = "Design"
	} else {
		fieldOfInterest = "All"
	}

	var level string
	if c.FormValue("beginner") == "Beginner" {
		level = "Beginner"
	} else if c.FormValue("intermediate") == "Intermediate" {
		level = "Intermediate"
	} else {
		level = "Advance"
	}

	userId, err := c.Cookie("userId")
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/register")
	}

	mentee := model.Mentee{
		ProfilePicture:  file.Filename,
		Language:        language,
		WorksAt:         worksAt,
		Phone:           phone,
		Designation:     designation,
		DateOfBirth:     dob,
		Location:        location,
		FieldOfInterest: fieldOfInterest,
		Level:           level,
		UserId:          userId.Value,
	}

	fmt.Println(mentee)
	result := model.DB.Create(&mentee)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(result)
	return c.Redirect(http.StatusTemporaryRedirect, "/mentee/profile")
}
