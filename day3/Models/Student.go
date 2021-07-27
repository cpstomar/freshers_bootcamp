package Models

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	FirstName		string 		`json:"firstname"`
	LastName		string 		`json:"lastname"`
	DOB   			string		`json:"DOB"`
	Address 		string 		`json:"address"`
}