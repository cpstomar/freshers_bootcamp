package Service

import (
	"store-api/Config"
	"store-api/Models"
	_ "github.com/go-sql-driver/mysql"
)

func Transaction(order *[]Models.Orders) (err error) {
	if err = Config.DB.Find(order).Error; err != nil {
		return err
	}
	return nil
}

