# NULL値の扱い

Gorm の[ドキュメントには次のようにある](https://gorm.io/ja_JP/docs/query.html#%E6%A7%8B%E9%80%A0%E4%BD%93-amp-%E3%83%9E%E3%83%83%E3%83%97%E3%81%A7%E3%81%AE%E6%9D%A1%E4%BB%B6%E6%8C%87%E5%AE%9A)。

```text
NOTE When querying with struct, GORM will only query with non-zero fields, that means if your field’s value is 0, '', false or other zero values, it won’t be used to build query conditions, for example:
要約：構造体を使ったクエリの場合、クエリに使えるのは0値でないフィールドだけ
```

つまり、空文字や0、nilは Where() や Updates() では無視される。

## NULL値を使用した検索

下記コード例では、Age や NameKanae が 0値なため、無視されていることがわかる。

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

  db.Where(&User{
    Name: "社員A",
  }).Find(&user)
  // => SELECT * FROM `users` WHERE `users`.`name` = '社員A'
}
```

空文字や0値で検索したい場合はポインタ型にする。以下はポインタ型にした例。conv パッケージは独自に定義したポインタ型に変換するパッケージ。

```go
package main

import (
  "sample/go-gorm-example/conv"
  "sample/go-gorm-example/infra"
)

type User struct {
  ID       uint
  Name     *string
  NameKana *string
  Age      *uint8
}

func main() {
  db := infra.DB().Debug()

  db.Exec("DROP TABLE users;")

  // User を AutoMigrate するだけで、Compnay も作られる
  db.AutoMigrate(&User{})

  db.Create(&User{
    Name:     conv.StrP("名前"),
    NameKana: conv.StrP("カナ"),
    Age:      conv.UbyteP(0),
  })
  // => INSERT INTO `users` (`name`,`name_kana`,`age`) VALUES ('名前','カナ',0)

  var user User
  db.Where(&User{Age: conv.UbyteP(0)}).First(&user)
  // => SELECT * FROM `users` WHERE `users`.`age` = 0 ORDER BY `users`.`id` LIMIT 1
}
```

## NULL値を使用した作成

以下のコードは NameKana に空文字を入れて Create した例である。

```go
package main

import (
  "sample/go-gorm-example/infra"
)

type User struct {
  ID       uint
  Name     string
  NameKana string
}

func main() {
  db := infra.DB().Debug()

  // User を AutoMigrate するだけで、Compnay も作られる
  db.AutoMigrate(&User{})

  db.Create(&User{
    Name: "社員",
  })
  // => INSERT INTO `users` (`name`,`name_kana`) VALUES ('社員','')
  db.Create(&User{
    Name:     "社員",
    NameKana: "",
  })
  //  INSERT INTO `users` (`name`,`name_kana`) VALUES ('社員','')
}
```

空文字が NameKana に代入されて INSERT INTO されている。Create の場合は空文字は無視されない。

下記のコードは Name と NameKana をポインタ型にして Create した例である。

```go
package main

import (
  "sample/go-gorm-example/conv"
  "sample/go-gorm-example/infra"
)

type User struct {
  ID       uint
  Name     *string
  NameKana *string
}

func main() {
  db := infra.DB().Debug()

  // User を AutoMigrate するだけで、Compnay も作られる
  db.AutoMigrate(&User{})

  db.Create(&User{
    Name: conv.StrP("社員"),
  })
  // => INSERT INTO `users` (`name`,`name_kana`) VALUES ('社員',NULL)

  db.Create(&User{
    Name:     conv.StrP("社員"),
    NameKana: conv.StrP(""),
  })
  // => INSERT INTO `users` (`name`,`name_kana`) VALUES ('社員','')

  db.Model(&User{}).
    Create(map[string]interface{}{
      "Name": "社員",
    })
  // => INSERT INTO `users` (`name`) VALUES ('社員')

  db.Model(&User{}).
    Create(map[string]interface{}{
      "Name":     "社員",
      "NameKana": "",
    })
  // => INSERT INTO `users` (`name`,`name_kana`) VALUES ('社員','')
}
```

空文字は空文字として扱われ、未指定(nil)の場合は無視される。
また、map[string]interface{}で指定した場合も同様に動作する。

下記コードは sql.NullString(NULL許容型) を使用した例である。ポインタ型と同様の動作となる。

```go
package main

import (
  "database/sql"
  "sample/go-gorm-example/infra"
)

type User struct {
  ID       uint
  Name     sql.NullString
  NameKana sql.NullString
}

func main() {
  db := infra.DB().Debug()

  // User を AutoMigrate するだけで、Compnay も作られる
  db.AutoMigrate(&User{})

  db.Create(&User{
    Name: sql.NullString{String: "社員", Valid: true},
  })
  // => INSERT INTO `users` (`name`,`name_kana`) VALUES ('社員',NULL)

  db.Create(&User{
    Name:     sql.NullString{String: "社員", Valid: true},
    NameKana: sql.NullString{String: "", Valid: false},
  })
  // => INSERT INTO `users` (`name`,`name_kana`) VALUES ('社員',NULL)
}
```

NULL許容型は次の通り

- [sql.NullBool](https://pkg.go.dev/database/sql#NullBool)
- [sql.NullByte](https://pkg.go.dev/database/sql#NullByte)
- [sql.NullFloat64](https://pkg.go.dev/database/sql#NullFloat64)
- [sql.NullInt16](https://pkg.go.dev/database/sql#NullInt16)
- [sql.NullInt32](https://pkg.go.dev/database/sql#NullInt32)
- [sql.NullInt64](https://pkg.go.dev/database/sql#NullInt64)
- [sql.NullString](https://pkg.go.dev/database/sql#NullString)
- [sql.NullTime](https://pkg.go.dev/database/sql#NullTime)

## NULL値を使用した更新

更新についても空文字や0、nilは無視される。

```go
package main

import (
  "sample/go-gorm-example/infra"
)

type User struct {
  ID       uint
  NameKana string
  Name     string
  Age      uint8
}

func main() {
  db := infra.DB().Debug()

  db.Exec("DROP TABLE users;")

  // User を AutoMigrate するだけで、Compnay も作られる
  db.AutoMigrate(&User{})

  db.Create(&User{
    Name:     "名前",
    NameKana: "カナ",
    Age:      10,
  })
  // => INSERT INTO `users` (`name_kana`,`name`,`age`) VALUES ('カナ','名前',10)

  var user User
  db.First(&user)
  // => SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

  user.Name = ""
  user.Age = 0
  db.Updates(&user)
  // => UPDATE `users` SET `name_kana`='カナ' WHERE `id` = 1
}
```

ポインタ型にすることで、空文字や0に更新することができる。

```go
package main

import (
  "sample/go-gorm-example/conv"
  "sample/go-gorm-example/infra"
)

type User struct {
  ID       uint
  NameKana *string
  Name     *string
  Age      *uint8
}

func main() {
  db := infra.DB().Debug()

  db.Exec("DROP TABLE users;")

  // User を AutoMigrate するだけで、Compnay も作られる
  db.AutoMigrate(&User{})

  db.Create(&User{
    Name:     conv.StrP("名前"),
    NameKana: conv.StrP("カナ"),
    Age:      conv.UbyteP(10),
  })
  // => INSERT INTO `users` (`name_kana`,`name`,`age`) VALUES ('カナ','名前',10)

  var user User
  db.First(&user)
  // => SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

  user.Name = conv.StrP("")
  user.Age = conv.UbyteP(0)
  db.Updates(&user)
  // => UPDATE `users` SET `name_kana`='カナ',`name`='',`age`=0 WHERE `id` = 1
}
```

## NULL値に更新する方法

値をNULL値にしたい場合は以下の３つの方法がある。構造体を使った更新ではSaveメソッド以外ではNULL値にできない。

- map[string]interface{}を使う
- gorm.Expr("NULL")を使う
- Save()メソッドを使う

下記は gorm.Expr("NULL") を使用した例。Update は単一のカラムのみ更新するメソッド。

```go
package main

import (
  "sample/go-gorm-example/conv"
  "sample/go-gorm-example/infra"

  "gorm.io/gorm"
)

type User struct {
  ID       uint
  NameKana *string
  Name     *string
  Age      *uint8
}

func main() {
  db := infra.DB().Debug()

  db.Exec("DROP TABLE users;")

  // User を AutoMigrate するだけで、Compnay も作られる
  db.AutoMigrate(&User{})

  db.Create(&User{
    Name:     conv.StrP("名前"),
    NameKana: conv.StrP("カナ"),
    Age:      conv.UbyteP(10),
  })
  // => INSERT INTO `users` (`name_kana`,`name`,`age`) VALUES ('カナ','名前',10)

  var user User
  db.First(&user)
  // => SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

  db.Model(&User{ID: user.ID}).Update("Name", gorm.Expr("NULL"))
  // => UPDATE `users` SET `name`=NULL WHERE `id` = 1

  db.Model(&User{ID: user.ID}).Updates(map[string]interface{}{ "Name": gorm.Expr("NULL") })
  // => UPDATE `users` SET `name`=NULL WHERE `id` = 1
}
```

下記は map[string]interface{} を使用した例。

```go
package main

import (
  "sample/go-gorm-example/conv"
  "sample/go-gorm-example/infra"

  "gorm.io/gorm"
)

type User struct {
  ID       uint
  NameKana *string
  Name     *string
  Age      *uint8
}

func main() {
  db := infra.DB().Debug()

  db.Exec("DROP TABLE users;")

  // User を AutoMigrate するだけで、Compnay も作られる
  db.AutoMigrate(&User{})

  db.Create(&User{
    Name:     conv.StrP("名前"),
    NameKana: conv.StrP("カナ"),
    Age:      conv.UbyteP(10),
  })
  // => INSERT INTO `users` (`name_kana`,`name`,`age`) VALUES ('カナ','名前',10)

  var user User
  db.First(&user)
  // => SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

  db.Model(&User{ID: user.ID}).Updates(map[string]interface{}{"Name": gorm.Expr("NULL")})
  // => UPDATE `users` SET `name`=NULL WHERE `id` = 1
}
```
