package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func NewDB(params ...string) {
	var err error
	dsn := DBUser + ":" + DBPassword + "@tcp(" + DBHost + ":" + DBPort + ")/" +
		DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}
}

func GetDBInstance() *gorm.DB {
	return DB
}
