package Service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"store-api/Config"
	"store-api/Models"
	"strconv"
)

func PlaceOrder(order *Models.Orders) (err error) {
	mu.Lock()
	var product Models.Product
	fmt.Println(strconv.FormatUint(uint64(order.ProductID),10))
	err = GetProductbyID(&product, strconv.FormatUint(uint64(order.ProductID), 10))
	if err != nil {
		panic(err)
	}
	product.Quantity -= 1
	//err = UpdateProduct(&product)
	Config.DB.Save(&product)
	if err != nil {
		panic(err)
	}
	err = Config.DB.Set("gorm:auto_preload",true).Create(order).Error;
	mu.Unlock()
	panic(err)
	if  err != nil {
		return err
	}
	return nil
}

func GetOrderByID(order *Models.Orders,id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *Models.User)(err error){
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}