package model

import "gorm.io/gorm"

type Mentor struct {
	gorm.Model
	Name            string
	ProfilePicture  string
	Email           string `gorm:"unique:true"`
	WorksAt         string
	Phone           string
	Designation     string
	FieldOfInterest string
	Password        string
}
