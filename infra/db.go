package infra

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USER = "root"
const DB_PASS = "pass"
const DB_HOST = "localhost"
const DB_PORT = "3306"
const DB_NAME = "go-gorm-example"
const DB_LOCA = "Local"

func DB() *gorm.DB {
	dsn := DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=" + DB_LOCA
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
