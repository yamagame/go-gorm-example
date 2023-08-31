# Gen

gen は定義したモデルから DAO (Data access object) をコード生成するパッケージ。

参考：[https://gorm.io/ja_JP/gen/](https://gorm.io/ja_JP/gen/)

以下、コード生成のサンプルコード。

```go
package main

import (
	"sample/go-gorm-example/infra"
	"sample/go-gorm-example/model"

	"gorm.io/gen"
)

// Dynamic SQL コメント部分のSQLを実行するメソッドを作成できる
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if age != 0}} AND age = @age{{end}}
	FilterWithNameAndAge(name, age uint) ([]gen.T, error)
}

func main() {
	db := infra.DB()

	g := gen.NewGenerator(gen.Config{
    // 生成ディレクトリ、パッケージ名になる
		OutPath: "./dao",
		// モード
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

  // gorm.DBを指定する
	g.UseDB(db)

	// model.User の DAO API を生成
	g.ApplyBasic(model.User{})

  // model.User と model.Company を使った Querier インタフェースで定義されたダイナミックSQLのタイプセーフAPIを生成
	g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})

	// コードを生成
	g.Execute()
}
```

生成した DAO の利用例。

```go
package main

import (
	"context"
	"fmt"
	"sample/go-gorm-example/conv"
	"sample/go-gorm-example/dao"
	"sample/go-gorm-example/infra"
	"sample/go-gorm-example/model"
)

func main() {
	ctx := context.TODO()
	db := infra.DB()

	query := dao.Use(db)
	u := query.WithContext(ctx).User

	// DAOを使用した User の作成
	u.Create(&model.User{
		Name: conv.StrP("社員"),
		Age:  conv.UintP(23),
	})

	// DAOを使用した User の検索
	user, err := u.First()
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
```

データベースからコードを生成するサンプル。

```go
package main

import (
	"sample/go-gorm-example/infra"

	"gorm.io/gen"
)

func main() {
	db := infra.DB()

	g := gen.NewGenerator(gen.Config{
		// 生成ディレクトリ、パッケージ名になる
		OutPath: "./infra/dao",
		// モード
		Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	// gorm.DBを指定する
	g.UseDB(db)

	g.ApplyBasic(
		// users テーブルから DAO を生成
		g.GenerateModel("users"),
	)
	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// コードを生成
	g.Execute()
}
```
