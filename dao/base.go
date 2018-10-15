package dao

import (
	Config "GoDemo/config"
	DbModel "GoDemo/model/db"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Init o
func Init() {
	db, err := gorm.Open("postgres", Config.ConnectionStr())
	if err != nil {
		panic("failed:" + err.Error())
	}
	defer db.Close()
	db.AutoMigrate(&DbModel.Members{})
	db.AutoMigrate(&DbModel.Orders{})
}

//Open o
func Open() *gorm.DB {
	db, err := gorm.Open("postgres", Config.ConnectionStr())
	if err != nil {
		panic("failed:" + err.Error())
	}
	return db
}
