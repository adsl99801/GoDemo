package dao

import (
	Config "GoDemo/config"
	DbModel "GoDemo/model/db"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//MembersDao o
type MembersDao struct {
}

//Init o
func (dao MembersDao) Init() {
	db, err := gorm.Open("postgres", Config.ConnectionStr())
	if err != nil {
		panic("failed:" + err.Error())
	}
	defer db.Close()
	db.AutoMigrate(&DbModel.Members{})
}

//All o
func (dao MembersDao) All() []DbModel.Members {
	db, err := gorm.Open("postgres", Config.ConnectionStr())
	if err != nil {
		panic("failed:" + err.Error())
	}
	defer db.Close()
	var members []DbModel.Members
	db.Where("Name is not null").Find(&members)
	return members
}

//Insert o
func (dao MembersDao) Insert() DbModel.Members {
	db, err := gorm.Open("postgres", Config.ConnectionStr())
	if err != nil {
		panic("failed:" + err.Error())
	}
	defer db.Close()
	member := DbModel.Members{Name: "test1", LOGINUID: "p1"}
	db.Create(&member)
	return member
}
