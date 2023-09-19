package main

import (
	"sample/go-gorm-example/infra"
	"sample/go-gorm-example/model"
)

type User struct {
	ID   uint
	Name string
}

func main() {
	db := infra.DB()

	// users テーブルの作成
	var err error
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.Company{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.Address{})
	if err != nil {
		panic(err)
	}

	// users テーブルに作成
	db.Create(&User{
		Name: "ユーザーA",
	})
}
