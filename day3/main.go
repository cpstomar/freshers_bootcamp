package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"hello/Config"
	"hello/Models"
	"hello/Routes"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.DropTable(&Models.Student{}, &Models.Marks{})
	Config.DB.AutoMigrate(&Models.Student{}, &Models.Marks{})
	r := Routes.SetupRouter()
	//running
	fmt.Println(Config.DB)
	r.Run()
}