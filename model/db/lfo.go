package db

import (
	"github.com/jinzhu/gorm"
)

//Order o
type Order struct {
	gorm.Model
	Title     string
	Contents  string
	FoldersID uint64
}

//Folders o
type Folders struct {
	gorm.Model
	MembersID uint64
}

//Model o
type Model struct {
	ID uint64 `gorm:"primary_key"`
}

//Members o
type Members struct {
	gorm.Model
	Name     string
	LOGINUID string
	Phone    string
}

//Tests o
type Tests struct {
	ID   int
	Name string
}
