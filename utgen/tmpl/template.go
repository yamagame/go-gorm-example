package tmpl

// QueryMethodTest query method test template
const QueryMethodTest = `
package dao

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dbName = "gen_test.db"

var db *gorm.DB
var once sync.Once

func init() {
	InitializeDB()
	db.AutoMigrate(&_another{})
}

const DB_USER = "root"
const DB_PASS = "pass"
const DB_HOST = "localhost"
const DB_PORT = "3306"
const DB_NAME = "go-gorm-example"

func InitializeDB() {
	once.Do(func() {
		var err error
		dsn := DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("open sqlite %q fail: %w", dbName, err))
		}
	})
}

func assert(t *testing.T, methodName string, res, exp interface{}) {
	if !reflect.DeepEqual(res, exp) {
		t.Errorf("%v() gotResult = %v, want %v", methodName, res, exp)
	}
}

type _another struct {
	ID uint64 ` + "`" + `gorm:"primaryKey"` + "`" + `
}

func (*_another) TableName() string { return "another_for_unit_test" }

func Test_Available(t *testing.T) {
	if !Use(db).Available() {
		t.Errorf("query.Available() == false")
	}
}

func Test_WithContext(t *testing.T) {
	query := Use(db)
	if !query.Available() {
		t.Errorf("query Use(db) fail: query.Available() == false")
	}

	type Content string
	var key, value Content = "gen_tag", "unit_test"
	qCtx := query.WithContext(context.WithValue(context.Background(), key, value))

	for _, ctx := range []context.Context{
		{{range $name,$d :=.Data -}}
		qCtx.{{$d.ModelStructName}}.UnderlyingDB().Statement.Context,
		{{end}}
	} {
		if v := ctx.Value(key); v != value {
			t.Errorf("get value from context fail, expect %q, got %q", value, v)
		}
	}
}

func Test_Transaction(t *testing.T) {
	query := Use(db)
	if !query.Available() {
		t.Errorf("query Use(db) fail: query.Available() == false")
	}

	err := query.Transaction(func(tx *Query) error { return nil })
	if err != nil {
		t.Errorf("query.Transaction execute fail: %s", err)
	}

	tx := query.Begin()

	err = tx.SavePoint("point")
	if err != nil {
		t.Errorf("query tx SavePoint fail: %s", err)
	}
	err = tx.RollbackTo("point")
	if err != nil {
		t.Errorf("query tx RollbackTo fail: %s", err)
	}
	err = tx.Commit()
	if err != nil {
		t.Errorf("query tx Commit fail: %s", err)
	}

	err = query.Begin().Rollback()
	if err != nil {
		t.Errorf("query tx Rollback fail: %s", err)
	}
}
`

// CRUDMethodTest CRUD method test
const CRUDMethodTest = `
package dao

import (
	"context"
	"fmt"
	"testing"

	"sample/go-gorm-example/infra/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&{{.Package}}.{{.ModelStructName}}{})
	if err != nil{
		fmt.Printf("Error: AutoMigrate(&{{.Package}}.{{.ModelStructName}}{}) fail: %s", err)
	}
}

func Test_{{.QueryStructName}}Query(t *testing.T) {
	{{.QueryStructName}} := new{{.ModelStructName}}(db)
	// {{.QueryStructName}} = *{{.QueryStructName}}.As({{.QueryStructName}}.TableName())
	_do := {{.QueryStructName}}.WithContext(context.Background()).Debug()

	primaryKey := field.NewString({{.QueryStructName}}.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <{{.TableName}}> fail:", err)
		return
	}
	
	_, ok := {{.QueryStructName}}.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from {{.QueryStructName}} success")
	}

	err = _do.Create(&{{.Package}}.{{.ModelStructName}}{})
	if err != nil {
		t.Error("create item in table <{{.TableName}}> fail:", err)
	}

	err = _do.Save(&{{.Package}}.{{.ModelStructName}}{})
	if err != nil {
		t.Error("create item in table <{{.TableName}}> fail:", err)
	}

	err = _do.CreateInBatches([]*{{.Package}}.{{.ModelStructName}}{ {}, {} }, 10)
	if err != nil {
		t.Error("create item in table <{{.TableName}}> fail:", err)
	}

	_, err = _do.Select({{.QueryStructName}}.ALL).Take()
	if err != nil {
		t.Error("Take() on table <{{.TableName}}> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <{{.TableName}}> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <{{.TableName}}> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <{{.TableName}}> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*{{.Package}}.{{.ModelStructName}}{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <{{.TableName}}> fail:", err)
	}

	_, err = _do.Select({{.QueryStructName}}.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <{{.TableName}}> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <{{.TableName}}> fail:", err)
	}

	_, err = _do.Select({{.QueryStructName}}.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <{{.TableName}}> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <{{.TableName}}> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <{{.TableName}}> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <{{.TableName}}> fail:", err)
	}
	
	_, err = _do.ScanByPage(&{{.Package}}.{{.ModelStructName}}{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <{{.TableName}}> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <{{.TableName}}> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <{{.TableName}}> fail:", err)
	}
	
	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <{{.TableName}}> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <{{.TableName}}> fail:", err)
	}
	
	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <{{.TableName}}> fail:", err)
	}
}
`
