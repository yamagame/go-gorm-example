package main

import (
	"sample/go-gorm-example/infra"
	"sample/go-gorm-example/model"

	"gorm.io/gen"
)

func main() {
	db := infra.DB()

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Company{})

	g := gen.NewGenerator(gen.Config{
		// 生成ディレクトリ、パッケージ名になる
		OutPath: "./infra/dao",
		// モード
		Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	// gorm.DBを指定する
	g.UseDB(db)

	// g.ApplyBasic(
	// 	// users テーブルから DAO を生成
	// 	g.GenerateModel("users"),
	// )
	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// コードを生成
	g.Execute()
}
