# Gormの基本

Gorm の主だった使い方を以下のまとめる。

参考：[GORMガイド](https://gorm.io/ja_JP/docs/index.html)

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

上記、Blogは下記のBlogと同じ定義となる。

```go
type Blog struct {
  ID    int64
  Name  string
  Email string
  Upvotes  int32
}
```

## Create / Take / Save / Delete

```go
package main

import (
  "fmt"
  "sample/go-gorm-example/infra"

  "gorm.io/gorm"
)

type User struct {
  gorm.Model
  Name  string
  Email string
}

func main() {
  db := infra.DB()

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

  // Save 更新
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

## golang-migrate/migrate

参考：https://github.com/golang-migrate/migrate
