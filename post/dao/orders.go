package dao

import (
	Config "GoDemo/config"
	DbModel "GoDemo/model/db"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//OrdersDao o
type OrdersDao struct {
}

//All o
func (dao OrdersDao) All() []DbModel.Orders {
	db, err := gorm.Open("postgres", Config.ConnectionStr())
	if err != nil {
		panic("failed:" + err.Error())
	}
	defer db.Close()
	var Orders []DbModel.Orders
	db.Where("Name is not null").Find(&Orders)
	return Orders
}
