package Service

import (
	"store-api/Config"
	"store-api/Models"
	_ "github.com/go-sql-driver/mysql"
)

func AddProduct(product *Models.Product) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductbyID(product *Models.Product,id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

func GetAllProduct(product *[]Models.Product) (err error) {
	if err = Config.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProduct(product *Models.Product) (err error) {
	mu.Lock()
	err = Config.DB.Update(product).Error;
	mu.Unlock()

	if err != nil {
		return err
	}
	return nil
}