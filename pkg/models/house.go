package models

import (
	"github.com/jinzhu/gorm"
	"houseAPI/pkg/config"
)



var db *gorm.DB
type House struct {
	gorm.Model
	Price           int  `json:"price"`
	SquareMeters    int  `json:"square_meters"`
	Floors          int  `json:"floors"`
	City            string `json:"city"`
	Region          string `json:"region"`
	ZipCode         int    `json:"zip_code"`
	ForSale         bool   `json:"for_sale"`
}

func init() {
	config.DBConnect()
	db = config.GetDB()
	db.AutoMigrate(&House{})
}

func (h *House) CreateHouse() *House{
	db.NewRecord(h)
	db.Create(&h)
	return h
}

func GetAllHouses() []House {
	var Houses []House
	db.Find(&Houses)
	return Houses
}

func GetHouseById(Id int64) (*House, *gorm.DB){
	var getHouse House
	db := db.Where("ID=?", Id).Find(&getHouse)
	return &getHouse, db
}

func DeleteHouse(ID int64) House{
	var house House
	db.Where("ID=?", ID).Delete(house)
	return house
}



