package main

import (
	"sample/go-gorm-example/infra"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func main() {
	db := infra.DB()

	g := gen.NewGenerator(gen.Config{
		// 生成ディレクトリ、パッケージ名になる
		OutPath: "./infra/dao",
		// モード
		Mode:              gen.WithoutContext | gen.WithQueryInterface,
		FieldWithIndexTag: true,
		FieldNullable:     true,
		FieldWithTypeTag:  true,
	})

	// gorm.DBを指定する
	g.UseDB(db)

	// 全てのテーブルを取得
	tableList, err := db.Migrator().GetTables()
	if err != nil {
		panic(err)
	}

	remove := func(l []string, item string) []string {
		for i, other := range l {
			if other == item {
				return append(l[:i], l[i+1:]...)
			}
		}
		return l
	}

	// 各テーブル毎にモデルを作成
	tables := make([]interface{}, len(tableList))
	company := g.GenerateModel("companies")
	user := g.GenerateModel("users",
		gen.FieldType("role", "*Role"),
		gen.FieldRelate(field.HasOne, "Company", company,
			&field.RelateConfig{
				RelatePointer: true,
				GORMTag:       field.GormTag{"foreignKey": []string{"CompanyID"}},
			},
		),
	)
	address := g.GenerateModel("addresses")

	tables = append(tables, user)
	tableList = remove(tableList, "users")
	tables = append(tables, company)
	tableList = remove(tableList, "companies")
	tables = append(tables, address)
	tableList = remove(tableList, "address")

	// 残りのテーブルのモデルを作成
	for _, tableName := range tableList {
		tables = append(tables, g.GenerateModel(tableName))
	}

	// DAOを生成
	g.ApplyBasic(
		tables...,
	)
	// コードを生成
	g.Execute()
}
