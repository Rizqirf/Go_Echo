package storage

// import (
// 	"log"

// 	config "github.com/Rizqirf/go_echo/config"
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/postgres"
// )

import (
	model "github.com/Rizqirf/go_echo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
  
var DB *gorm.DB


func DbConnect() {
	dsn := "host=localhost user=postgres password=password dbname=go_crud port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err !=nil {
		panic("Connection to Database Failed!!")
	}

	err = db.AutoMigrate(&model.PropertyPromotion{})

	if err != nil {
		return
	}

	DB = db
}

func GetDBInstance() *gorm.DB {
	return DB
}