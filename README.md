# Gormの基本

Gorm の主だった使い方を以下のまとめる。

参考：[GORMガイド](https://gorm.io/ja_JP/docs/index.html)

## メソッド一覧

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

## モデルの定義

以下はモデルの定義例。それぞれの変数がスネークケースでテーブルのフィールドになる。例えば「MemberNumber」は「member_number」となる。

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

この構造体を使ってレコードを作成すると users テーブルにレコードを作成できる。テーブル名は自動的に複数形の名前になる。

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

インデックスは「`gorm:"index"`」を使用する

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

次の users テーブルを例に記述する。

```go
type User struct {
  ID uint
  Name string
  NameKana string
  Age int
}
```

プライマリーキーを指定して検索する場合。

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
    Limit(3).             // 指定件数のみ
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

更新は Where メソッドを指定しないければデフォルトではエラーになる。

参考：[Global Updatesを防ぐ](https://gorm.io/ja_JP/docs/update.html#Global-Updates%E3%82%92%E9%98%B2%E3%81%90)

### Save

すべてのカラムを更新する場合は Save を使用する。すべて更新されるため、使用しない方が無難。

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

一つのカラムを更新する場合は Update を使用する。単一のカラムだけしか更新できないため用途限定。

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
  db.Model(&user).Select("Name", "Age").Updates(User{Name: "new_name", Age: 20})

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
func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    q := r.URL.Query()
    page, _ := strconv.Atoi(q.Get("page"))
    if page <= 0 {
      page = 1
    }

    pageSize, _ := strconv.Atoi(q.Get("page_size"))
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

db.Scopes(Paginate(r)).Find(&users)
db.Scopes(Paginate(r)).Find(&articles)
```

## NULLの扱い

```go
type User struct {
  ID       uint
  Name     string
  NameKana string
  Age      int
}
```

上記のような構造体があった場合、NameやAgeに0値が代入されると Where() や Updates() では無視される。

```go
package main

import (
  "sample/go-gorm-example/infra"
)

type User struct {
  ID       uint
  Name     string
  NameKana string
  Age      uint8
}

func main() {
  db := infra.DB().Debug()

  db.AutoMigrate(&User{})

  var user User
  db.Where(&User{
    Name:     "社員A",
    NameKana: "",
    Age:      0,
  }).Find(&user)
  // => SELECT * FROM `users` WHERE `users`.`name` = '社員A'
}
```

上記例では WHERE 句に Name はあるが、NameKana や Age は指定がない。これは Gorm によって0値が無視されたため。

Gorm の[ドキュメントには次のようにある](https://gorm.io/ja_JP/docs/query.html#%E6%A7%8B%E9%80%A0%E4%BD%93-amp-%E3%83%9E%E3%83%83%E3%83%97%E3%81%A7%E3%81%AE%E6%9D%A1%E4%BB%B6%E6%8C%87%E5%AE%9A)。

```text
NOTE When querying with struct, GORM will only query with non-zero fields, that means if your field’s value is 0, '', false or other zero values, it won’t be used to build query conditions, for example:
要約：クエリに使えるのは0値でないフィールドだけ
```

この仕様はフィールドを空文字にしたり、NULL値にしたくでも通常はできないことを意味している。

フィールドを0値にする場合は以下の２つの方法がある。

- フィールドをポインタ型にする
- [database/sql](https://pkg.go.dev/database/sql) パッケージが用意したNULL許容型を使用する

NULL許容型は次の通り

- [sql.NullBool](https://pkg.go.dev/database/sql#NullBool)
- [sql.NullByte](https://pkg.go.dev/database/sql#NullByte)
- [sql.NullFloat64](https://pkg.go.dev/database/sql#NullFloat64)
- [sql.NullInt16](https://pkg.go.dev/database/sql#NullInt16)
- [sql.NullInt32](https://pkg.go.dev/database/sql#NullInt32)
- [sql.NullInt64](https://pkg.go.dev/database/sql#NullInt64)
- [sql.NullString](https://pkg.go.dev/database/sql#NullString)
- [sql.NullTime](https://pkg.go.dev/database/sql#NullTime)

```go
// ポインタ型にした場合
type User struct {
  ID       uint
  Name     *string
  NameKana *string
  Age      *uint8
}

// NULL許容型にした場合
type User struct {
  ID       uint
  Name     sql.NullString
  NameKana sql.NullString
  Age      sql.NullByte
}
```

それぞれについてインスタンス化する場合は以下のようになる。conv パッケージは独自に定義したポインタ型に変換するパッケージ。

```go
// ポインタ型にした場合
user := User{
  Name: conv.StrP("名前"),
  NameKana: conv.StrP("カナ"),
  Age: conv.UbyteP(10),
}

// NULL許容型にした場合
user := User{
  Name: sql.NullString{"名前", true},
  NameKana: sql.NullString{"カナ", true},
  Age: sql.NullByte{10, true},
}
```

NULLを代入する場合は次のようにする。

```go
// ポインタ型にした場合
user := User{
  Name: nil,  // 省略可
  NameKana: nil,  // 省略可
  Age: nil,  // 省略可
}

// NULL許容型にした場合
user := User{
  Name: sql.NullString{"", false},
  NameKana: sql.NullString{"", false},
  Age: sql.NullByte{0, false},
}
```

値がNULLかどうかは次のように判定する。

```go
// ポインタ型にした場合
if user.Name != nil {
  // 値あり
}

// NULL許容型にした場合
if user.Name.Valid {
  // 値あり
}
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

## ER図

参考：[ER図（Entity Relationship Diagram）](https://segakuin.com/diagram/entity-relationship-diagram.html)

## golang-migrate/migrate

参考：https://github.com/golang-migrate/migrate
