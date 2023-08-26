# いけないコピペプログラミング

下記のコードは User テーブルと Company テーブルを CSV ファイルにする例である。
WriteUserCSV が User テーブルをCSV化する関数、WriteCompanyCSV が Company テーブルをCSV化する関数である。それぞれのテーブル専用の関数が実装されている。

```go
// 例1
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sample/go-gorm-example/infra"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	NameKana string
	Age      uint
}

// User モデルをCSV化する関数
func WriteUserCSV(db *gorm.DB, w io.Writer) {
	var users []User
	db.Find(&users)

	cells := [][]string{}
	cells = append(cells, []string{
		"名前",
		"カナ",
		"年齢",
	})
	for _, user := range users {
		record := []string{}
		record = append(record, user.Name)
		record = append(record, user.NameKana)
		record = append(record, fmt.Sprint(user.Age))
		cells = append(cells, record)
	}

	writer := csv.NewWriter(w)
	writer.WriteAll(cells)
}

type Company struct {
	gorm.Model
	Name     string
	NameKana string
	Capital  uint
}

// Company モデルをCSV化する関数
func WriteCompanyCSV(db *gorm.DB, w io.Writer) {
	var companyies []Company
	db.Find(&companyies)

	cells := [][]string{}
	cells = append(cells, []string{
		"名前",
		"カナ",
		"資本金",
	})
	for _, company := range companyies {
		record := []string{}
		record = append(record, company.Name)
		record = append(record, company.NameKana)
		record = append(record, fmt.Sprint(company.Capital))
		cells = append(cells, record)
	}

	writer := csv.NewWriter(w)
	writer.WriteAll(cells)
}

func main() {
	db := infra.DB()
  WriteUserCSV(db, os.Stdout)
	WriteCompanyCSV(db, os.Stdout)
}
```

次の例は、CSV ファイルを生成する責務を WriteCSV に持たせ、User と Company には CSV 化するために必要な情報だけを持たせている。
Generics や Interface を使用することで、データ定義と処理を分離できる。

参考：[Go generics で指定された型のオブジェクトをインターフェイスとして返す](https://qiita.com/Nabetani/items/f9aa0bb668d471e39186)

```go
// 例2
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sample/go-gorm-example/infra"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	NameKana string
	Age      uint
}

func (User) Header() []string {
	return []string{"名前",
		"カナ",
		"年齢",
	}
}

func (x *User) Record() []string {
	record := []string{}
	record = append(record, x.Name)
	record = append(record, x.NameKana)
	record = append(record, fmt.Sprint(x.Age))
	return record
}

type Company struct {
	gorm.Model
	Name     string
	NameKana string
	Capital  uint
}

func (Company) Header() []string {
	return []string{
		"名前",
		"カナ",
		"資本金",
	}
}

func (x *Company) Record() []string {
	record := []string{}
	record = append(record, x.Name)
	record = append(record, x.NameKana)
	record = append(record, fmt.Sprint(x.Capital))
	return record
}

type CSVRecordInterface interface {
	Header() []string
	Record() []string
}

type CSVRecordConstraint[T any] interface {
	CSVRecordInterface
	*T
}

// User も Company もこの関数で CSV 化する
func WriteCSV[T any, PT CSVRecordConstraint[T]](db *gorm.DB, w io.Writer) {
	var v T
	pv := PT(&v)
	var rows []PT
	db.Find(&rows)

	cells := [][]string{}
	cells = append(cells, pv.Header())
	for _, row := range rows {
		cells = append(cells, row.Record())
	}

	writer := csv.NewWriter(w)
	writer.WriteAll(cells)
}

func main() {
	db := infra.DB()

	WriteCSV[User](db, os.Stdout)
	WriteCSV[Company](db, os.Stdout)
}
```

例1の方がコピペ後、異なる部分を変更するだけでよいため一見開発が簡単である。しかし、CSVファイルの作成処理に変更が加わった時、すべての作成関数に手を入れる必要があり、テーブルが増えたときに大変になる。  
例2であれば、データ定義と処理を分離できているため WriteCSV のみの変更で済む。

例1のパターンをここでは「コピペプログラミング」と呼んでいる。後々、保守が大変になるため、極力「コピペプログラミング」は避けなければいけない。
