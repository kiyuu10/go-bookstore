package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// mysql://avnadmin:AVNS_sXLm_r_yre9iqhKDvYJ@huymysql-htl-46e7.a.aivencloud.com:15121/defaultdb?ssl-mode=REQUIRED
func Connect() {
	//d, err := gorm.Open("mysql",
	//	"akhil:Axlesharma@12@/simplerest?charset=utf8&parseTime=True&loc=Local")
	d, err := gorm.Open("mysql",
		"mysql://avnadmin:AVNS_sXLm_r_yre9iqhKDvYJ@huymysql-htl-46e7.a.aivencloud.com:15121/defaultdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
