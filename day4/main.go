package main

import (
	"fmt"
	"store-api/Config"
	"store-api/Models"
	"store-api/Router"
	"github.com/jinzhu/gorm"
)

var err error
func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.DropTable(&Models.Orders{}, &Models.Product{}, &Models.User{})
	Config.DB.AutoMigrate(&Models.Orders{}, &Models.Product{}, &Models.User{})
	r := Router.SetupRouter()
	//running
	fmt.Println(Config.DB)
	r.Run()
}
