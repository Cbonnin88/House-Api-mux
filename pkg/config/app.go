package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)
var (
	db *gorm.DB
)

func DBConnect () {
	// pwd isn't actual password

	dbConnect, err := gorm.Open("mysql", "mysql","root:password@(127.0.0.1:3306)/houses_db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println("Error: unable to connect to the database")
	}
	db = dbConnect

	fmt.Println("DataBase connection successfully Opened")

}

func GetDB() *gorm.DB{
	return db
}




