package Models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name		string	`json:"name"`
	Price		int		`json:"price"`
	Quantity	int		`json:"quantity"`
}