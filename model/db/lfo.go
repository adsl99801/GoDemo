package db

import (
	"github.com/jinzhu/gorm"
)

//Orders o
type Orders struct {
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

//LoginHistorys o
type LoginHistorys struct {
	gorm.Model
	MembersID string
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
