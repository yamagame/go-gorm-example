# Gormの基本

Gorm の主だった使い方を以下のまとめる。

参考：[GORMガイド](https://gorm.io/ja_JP/docs/index.html)

## モデルの定義

以下はモデルの定義例。この構造体は users テーブルになる。テーブル名は自動的に複数形の名前になる。
それぞれの変数がスネークケースでテーブルのフィールドになる。例えば「MemberNumber」は「member_number」となる。

```go
type User struct {
  ID           uint
  Name         string
  Email        *string
  Age          uint8
  Birthday     *time.Time
  MemberNumber sql.NullString
  ActivatedAt  sql.NullTime
  CreatedAt    time.Time
  UpdatedAt    time.Time
}
```

以下、Userレコードを一件作成する例。

```go
package main

import (
  "sample/go-gorm-example/infra"
)

type User struct {
  ID   uint
  Name string
}

func main() {
  db := infra.DB()

  // users テーブルの作成
  db.AutoMigrate(&User{})

  // users テーブルに作成
  db.Create(&User{
    Name: "ユーザーA",
  })
}
```

プライマリキーはデフォルトでは「ID」であるが、タグを使用することで変更することもできる。

```go
type User struct {
  UserKey      uint      `gorm:"primaryKey"`
  Name         string
           :
           :
}
```

インデックスを作成する場合は「`gorm:"index"`」を使用する

```go
type User struct {
  　　　　　:
  Age          uint8     `gorm:"index"`
  　　　　　:
}
```

テーブルのフィールドとして扱いたくない場合は「`gorm:"-:all"`」を使用する。

```go
type User struct {
          :
  FullEmail    *string   `gorm:"-:all"`
          :
}
```

参考：[フィールドレベルの権限](https://gorm.io/ja_JP/docs/models.html#%E3%83%95%E3%82%A3%E3%83%BC%E3%83%AB%E3%83%89%E3%83%AC%E3%83%99%E3%83%AB%E3%81%AE%E6%A8%A9%E9%99%90)

「ID」「CreatedAt」「UpdatedAt」「DeletedAt」は「gorm.Model」を継承し定義することもできる。

```go
type User struct {
  gorm.Model               // この埋め込みで「ID」「CreatedAt」「UpdatedAt」「DeletedAt」が定義される。
  Name         string
  Email        *string
  Age          uint8
  Birthday     *time.Time
  MemberNumber sql.NullString
  ActivatedAt  sql.NullTime
}
```

gorm.Model の定義は以下の通り。

```go
// gorm.Modelの定義
type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}
```

「`gorm:"embedded"`」タグを使用して構造体を埋め込むことができる。

```go
type Author struct {
  Name  string
  Email string
}

type Blog struct {
  ID      int
  Author  Author `gorm:"embedded"`
  Upvotes int32
}
```

上記、Blog は下記の Blog と同じ定義となる。

```go
type Blog struct {
  ID    int64
  Name  string
  Email string
  Upvotes  int32
}
```

## Create / Take / Save / Delete

以下、Create、Take、Save、Delete の例

```go
package main

import (
  "fmt"
  "sample/go-gorm-example/infra"

  "gorm.io/gorm"
)

// User 構造体で users テーブルのテーブル定義
type User struct {
  gorm.Model
  Name  string
  Email string
}

func main() {
  db := infra.DB()

  // users テーブルの作成
  db.AutoMigrate(&User{})

  var err error

  // Create 作成
  err = db.Create(&User{
    Name:  "ユーザーA",
    Email: "user-a@sampele.com",
  }).Error
  if err != nil {
    panic(err)
  }

  // Take 一つ取得
  var user User
  err = db.Take(&user).Error
  if err != nil {
    panic(err)
  }
  fmt.Println(user)

  // Save はすべてのフィールドを更新する
  user.Name = "ユーザーA-2"
  err = db.Save(&user).Error
  if err != nil {
    panic(err)
  }
  fmt.Println(user)

  // Delete 削除
  err = db.Delete(&user).Error
  if err != nil {
    panic(err)
  }
}
```

## クエリ

下記コードはプライマリーキーを指定して検索する場合。

```go
package main

import (
  "sample/go-gorm-example/infra"
)

type User struct {
  ID       uint
  Name     string
  NameKana string
  Age      int
}

func main() {
  db := infra.DB()

  db.AutoMigrate(&User{})

  // プライマリーキーを入れて検索
  user := User{ID: 3}
  db.First(&user)

  // First() でプライマリーキーを指定
  var result User
  db.First(&result, 3)
}
```

テーブルに定義されたフィールドのうち必要なフィールドのみ取り出す場合は Model() を使用する。Name だけ取得するコードが下記になる。

```go
package main

import (
  "sample/go-gorm-example/infra"
)

type User struct {
  ID       uint
  Name     string
  NameKana string
  Age      int
}

type UserName struct {
  Name string
}

func main() {
  db := infra.DB()

  db.AutoMigrate(&User{})

  var user UserName
  // users テーブルから Name だけ取得
  db.Model(&User{}).Take(&user)

  // Pluck を使用して names に取得
  var names []string
  db.Model(&User{}).
    Limit(3).             // 3件まで
    Pluck("Name", &names) // Pluck は一つのカラムのみ取得できる
}
```

以下、Where() と Find() を使用したクエリの例

```go
package main

import (
  "sample/go-gorm-example/infra"
)

type User struct {
  ID       uint
  Name     string
  NameKana string
  Age      int
}

func main() {
  db := infra.DB()

  db.AutoMigrate(&User{})

  for i := 5; i < 15; i++ {
    db.Create(&User{
      Name: "name",
      Age:  i,
    })
  }

  var users []User
  // Age が 10 以上のユーサーを検索
  db.Where("Age >= ?", 10).Find(&users)

  // Age が 10 のユーザーを検索
  db.Where(&User{Age: 10}).Find(&users)
}
```

参考：[レコードの取得](https://gorm.io/ja_JP/docs/query.html#%E6%96%87%E5%AD%97%E5%88%97%E6%9D%A1%E4%BB%B6)  
参考：[高度なクエリ](https://gorm.io/ja_JP/docs/advanced_query.html)

## 更新

更新は Where メソッドを指定しないければ、デフォルトではエラーになる。

参考：[Global Updatesを防ぐ](https://gorm.io/ja_JP/docs/update.html#Global-Updates%E3%82%92%E9%98%B2%E3%81%90)

### Save

すべてのカラムを更新する場合は Save を使用する。すべてのフィールドが更新される。

```go
package main

import (
  "sample/go-gorm-example/infra"
)

type User struct {
  ID       uint
  Name     string
  NameKana string
  Age      int
}

func main() {
  db := infra.DB().Debug()

  db.AutoMigrate(&User{})

  // すべてのカラムを更新、プライマリーキーがなければ新規に作成される
  // NameKana には空文字が保存される
  db.Save(&User{
    Name: "名前A",
    Age:  20,
  })
}
```

### Update

単一のカラムを更新する場合は Update を使用する。

```go
package main

import (
  "sample/go-gorm-example/infra"
)

type User struct {
  ID       uint
  Name     string
  NameKana string
  Age      int
}

func main() {
  db := infra.DB()

  db.AutoMigrate(&User{})

  // Age が 10 以上のユーザーの Name を new-name に変更
  db.Model(&User{}).Where("Age >= ?", 10).Update("name", "new-name")

  // Find() で検索
  var result []User
  db.Limit(10).Find(&result)
}
```

### Updates

複数のカラムを更新する場合は Updates を使用する。

```go
package main

import (
  "sample/go-gorm-example/infra"
)

type User struct {
  ID       uint
  Name     string
  NameKana string
  Age      int
}

func main() {
  db := infra.DB()

  db.AutoMigrate(&User{})

  // Name と Age を更新
  user := User{ID: 3}
  db.Model(&user).Select("Name", "Age").
     Updates(User{
      Name: "new_name",
      Age: 20,
     })

  // First() でプライマリーキーを指定
  var result User
  err := db.First(&result, 3).Error
  if errors.Is(err, gorm.ErrRecordNotFound) {
    // レコードが見つからなかった
  }
}
```

## ページング

参考：[ページング](https://gorm.io/ja_JP/docs/scopes.html#%E3%83%9A%E3%83%BC%E3%82%B8%E3%83%B3%E3%82%B0)

```go
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    if page <= 0 {
      page = 1
    }

    switch {
    case pageSize > 100:
      pageSize = 100
    case pageSize <= 0:
      pageSize = 10
    }

    offset := (page - 1) * pageSize
    return db.Offset(offset).Limit(pageSize)
  }
}

// 1ページ10件として1ページ目を取得
db.Scopes(Paginate(1, 10)).Find(&users)

// 1ページ100件として2ページ目を取得
db.Scopes(Paginate(2, 100)).Find(&articles)
```

## アソシエーション

参考：[アソシエーション](https://gorm.io/ja_JP/docs/belongs_to.html)

次のような構造体を定義すると、User に Compnay を所属させることができる。

```go
type Company struct {
  gorm.Model
  Code string
  Name string
}

type User struct {
  gorm.Model
  Name string
  CompanyID int
  Company Company
}
```

Company の外部キーはデフォルトでは ComnapyID になる。外部キーが存在しない場合はエラーになる。
外部キーを変更したい場合は次のようなタグを付け加える。

```go
type User struct {
  gorm.Model
  Name string
  CompanyRefer int
  Company Company `gorm:"foreignKey:CompanyRefer"`
  // Company の外部キーが CompanyRefer になる
}
```

Company の参照キーは通常はプライマリキーになる。
Compnay 側の参照キーを ID から Code に変更したい場合は次のようにタグを付け加える。

```go
type User struct {
  gorm.Model
  Name string
  CompanyID int
  Company Company `gorm:"references:Code"`
  // Company の参照キーが ID から Code になる
}
```

User 構造体を AutoMigrate すると関連テーブルも一緒に作成される。

```go
package main

import (
  "sample/go-gorm-example/infra"

  "gorm.io/gorm"
)

type Company struct {
  gorm.Model
  Code string
  Name string
}

type User struct {
  gorm.Model
  Name      string
  CompanyID uint
  Company   Company
}

func main() {
  db := infra.DB().Debug()

  // User を AutoMigrate するだけで、Compnay も作られる
  db.AutoMigrate(&User{})

  // 会社Aに所属の社員Aを作成
  db.Create(&User{
    Name: "社員A",
    Company: Company{
      Code: "001",
      Name: "会社A",
    },
  })

  // Compnay も含めて社員Aの情報を取得
  var user User
  db.Preload("Company").Where(&User{Name: "社員A"}).Find(&user)

  // または
  db.Preload(clause.Associations).Where(&User{Name: "社員A"}).Find(&user)
}
```

Preload または Joins を使うことで、belongs to リレーションの [Eager Loading](https://gorm.io/ja_JP/docs/preload.html) を行うことができる。

## トランザクション

参考：[トランザクション](https://gorm.io/ja_JP/docs/transactions.html)

```go
db.Transaction(func(tx *gorm.DB) error {
  // トランザクション内でのデータベース処理を行う(ここでは `db` ではなく `tx` を利用する)
  if err := tx.Create(&Animal{Name: "Giraffe"}).Error; err != nil {
    // 何らかのエラーを返却するとロールバックされる
    return err
  }

  if err := tx.Create(&Animal{Name: "Lion"}).Error; err != nil {
    return err
  }

  // nilが返却されるとトランザクション内の全処理がコミットされる
  return nil
})
```

## NULL値の扱い

Gorm ではフィールドが0値(NULLや0、空文字)の場合は検索やアップデートのフィールドとして無視される。
フィールドの型は初期値の代入でリテラルが使えない点を除き、ポインタ型が便利な印象。

「[NULL値の扱い](./NULLVAL.md)」を参照

## リファレンス

参考:[https://pkg.go.dev/github.com/jinzhu/gorm](https://pkg.go.dev/github.com/jinzhu/gorm)

- [AutoMigrate](https://gorm.io/ja_JP/docs/migration.html#Auto-Migration)
- [Debug](https://gorm.io/ja_JP/docs/logger.html#%E3%83%87%E3%83%90%E3%83%83%E3%82%B0)
- [Model](https://gorm.io/ja_JP/docs/advanced_query.html#%E5%8F%96%E5%BE%97%E7%B5%90%E6%9E%9C%E3%82%92%E3%83%9E%E3%83%83%E3%83%97%E3%81%AB%E4%BB%A3%E5%85%A5)
- [Table](https://gorm.io/ja_JP/docs/advanced_query.html#%E5%8F%96%E5%BE%97%E7%B5%90%E6%9E%9C%E3%82%92%E3%83%9E%E3%83%83%E3%83%97%E3%81%AB%E4%BB%A3%E5%85%A5)
- [Select](https://gorm.io/ja_JP/docs/update.html#%E9%81%B8%E6%8A%9E%E3%81%97%E3%81%9F%E3%83%95%E3%82%A3%E3%83%BC%E3%83%AB%E3%83%89%E3%82%92%E6%9B%B4%E6%96%B0%E3%81%99%E3%82%8B)
- [Where](https://gorm.io/ja_JP/docs/advanced_query.html#%E3%82%B5%E3%83%96%E3%82%AF%E3%82%A8%E3%83%AA)
- [Omit](https://gorm.io/ja_JP/docs/create.html#%E9%AB%98%E5%BA%A6%E3%81%AA%E6%A9%9F%E8%83%BD)
- [Offset](https://gorm.io/ja_JP/docs/query.html#Limit-amp-Offset)
- [Limit](https://gorm.io/ja_JP/docs/query.html#Limit-amp-Offset)
- [Count](https://gorm.io/ja_JP/docs/advanced_query.html#Count)
- [Preload](https://gorm.io/ja_JP/docs/preload.html)
- [Joins](https://gorm.io/ja_JP/docs/preload.html#Joins-%E3%81%AB%E3%82%88%E3%82%8B-Preloading)
- [Not](https://gorm.io/ja_JP/docs/query.html#Not%E6%9D%A1%E4%BB%B6)
- [Or](https://gorm.io/ja_JP/docs/query.html#OR%E6%9D%A1%E4%BB%B6)
- [Create](https://gorm.io/ja_JP/docs/create.html#%E3%83%AC%E3%82%B3%E3%83%BC%E3%83%89%E3%82%92%E4%BD%9C%E6%88%90%E3%81%99%E3%82%8B)
- [Save](https://gorm.io/ja_JP/docs/update.html#%E3%81%99%E3%81%B9%E3%81%A6%E3%81%AE%E3%83%95%E3%82%A3%E3%83%BC%E3%83%AB%E3%83%89%E3%82%92%E4%BF%9D%E5%AD%98%E3%81%99%E3%82%8B)
- [Find](https://gorm.io/ja_JP/docs/query.html#%E3%81%99%E3%81%B9%E3%81%A6%E3%81%AE%E3%82%AA%E3%83%96%E3%82%B8%E3%82%A7%E3%82%AF%E3%83%88%E3%82%92%E5%8F%96%E5%BE%97%E3%81%99%E3%82%8B)
- [Scan](https://gorm.io/ja_JP/docs/query.html#Scan)
- [First](https://gorm.io/ja_JP/docs/query.html#%E5%8D%98%E4%B8%80%E3%81%AE%E3%82%AA%E3%83%96%E3%82%B8%E3%82%A7%E3%82%AF%E3%83%88%E3%82%92%E5%8F%96%E5%BE%97%E3%81%99%E3%82%8B)
- [Last](https://gorm.io/ja_JP/docs/query.html#%E5%8D%98%E4%B8%80%E3%81%AE%E3%82%AA%E3%83%96%E3%82%B8%E3%82%A7%E3%82%AF%E3%83%88%E3%82%92%E5%8F%96%E5%BE%97%E3%81%99%E3%82%8B)
- [Take](https://gorm.io/ja_JP/docs/query.html#%E5%8D%98%E4%B8%80%E3%81%AE%E3%82%AA%E3%83%96%E3%82%B8%E3%82%A7%E3%82%AF%E3%83%88%E3%82%92%E5%8F%96%E5%BE%97%E3%81%99%E3%82%8B)
- [Pluck](https://gorm.io/ja_JP/docs/advanced_query.html#Pluck)
- [Update](https://gorm.io/ja_JP/docs/update.html#%E5%8D%98%E4%B8%80%E3%81%AE%E3%82%AB%E3%83%A9%E3%83%A0%E3%82%92%E6%9B%B4%E6%96%B0%E3%81%99%E3%82%8B)
- [Updates](https://gorm.io/ja_JP/docs/update.html#%E8%A4%87%E6%95%B0%E3%81%AE%E3%82%AB%E3%83%A9%E3%83%A0%E3%82%92%E6%9B%B4%E6%96%B0%E3%81%99%E3%82%8B)
- [Delete](https://gorm.io/ja_JP/docs/delete.html#%E3%83%AC%E3%82%B3%E3%83%BC%E3%83%89%E3%82%92%E5%89%8A%E9%99%A4)
- [Raw](https://gorm.io/ja_JP/docs/sql_builder.html#Row-amp-Rows)
- [Raws](https://gorm.io/ja_JP/docs/advanced_query.html#Iteration)
- [Close](https://gorm.io/ja_JP/docs/advanced_query.html#Iteration)
- [Next](https://gorm.io/ja_JP/docs/advanced_query.html#Iteration)
- [ScanRows](https://gorm.io/ja_JP/docs/advanced_query.html#Iteration)
- [FindInBatches](https://gorm.io/ja_JP/docs/advanced_query.html#FindInBatches)
- [Exec](https://gorm.io/ja_JP/docs/sql_builder.html#Raw-SQL)
- [ToSQL](https://gorm.io/ja_JP/docs/sql_builder.html#ToSQL)
- [Connection](https://gorm.io/ja_JP/docs/sql_builder.html#Connection)
- [Session](https://gorm.io/ja_JP/docs/session.html)
- [Clauses](https://gorm.io/ja_JP/docs/sql_builder.html#Clauses)
- [Association](https://gorm.io/ja_JP/docs/associations.html#Association-Mode)
- [Append](https://gorm.io/ja_JP/docs/associations.html#Association-Mode)
- [Replace](https://gorm.io/ja_JP/docs/associations.html#Association-Mode)
- [Scopes](https://gorm.io/ja_JP/docs/advanced_query.html#Scopes)

## Gen

gen を使用すると DAO (Data access object) をコード生成できる。

「[Gen](./CODEGEN.md)」を参照

## ER図

参考：[ER図（Entity Relationship Diagram）](https://segakuin.com/diagram/entity-relationship-diagram.html)

## golang-migrate/migrate

参考：https://github.com/golang-migrate/migrate
