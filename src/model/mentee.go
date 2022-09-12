package model

import "gorm.io/gorm"

type Mentee struct {
	gorm.Model
	ProfilePicture  string
	Language        string
	WorksAt         string
	Phone           string
	Designation     string
	DateOfBirth     string
	Location        string
	FieldOfInterest string
	Level           string
	UserId          string
	User            User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
