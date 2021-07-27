package Models

import "github.com/jinzhu/gorm"

type Orders struct {
	gorm.Model
	OrderID		uint64 		`json:"orderID"`
	User 		User		`json:"user" binding:"required" gorm:"foreignkey:UserID"`
	UserID		uint64
	Product 	Product		`json:"product" binding:"required" gorm:"foreignkey:ProductID"`
	ProductID   uint64
	Status 		string 		`json:"msg"`
}
