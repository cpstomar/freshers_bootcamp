package Models

import "github.com/jinzhu/gorm"

type Marks struct {
	gorm.Model
	Student   Student `json:"student" gorm:"foreignkey:StudentID"`
	StudentID uint64
	Marks     int
	Subject   string
}