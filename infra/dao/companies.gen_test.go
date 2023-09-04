
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
	err := db.AutoMigrate(&model.Company{})
	if err != nil{
		fmt.Printf("Error: AutoMigrate(&model.Company{}) fail: %s", err)
	}
}

func Test_companyQuery(t *testing.T) {
	company := newCompany(db)
	// company = *company.As(company.TableName())
	_do := company.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(company.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <companies> fail:", err)
		return
	}
	
	_, ok := company.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from company success")
	}

	err = _do.Create(&model.Company{})
	if err != nil {
		t.Error("create item in table <companies> fail:", err)
	}

	err = _do.Save(&model.Company{})
	if err != nil {
		t.Error("create item in table <companies> fail:", err)
	}

	err = _do.CreateInBatches([]*model.Company{ {}, {} }, 10)
	if err != nil {
		t.Error("create item in table <companies> fail:", err)
	}

	_, err = _do.Select(company.ALL).Take()
	if err != nil {
		t.Error("Take() on table <companies> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <companies> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <companies> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <companies> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.Company{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <companies> fail:", err)
	}

	_, err = _do.Select(company.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <companies> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <companies> fail:", err)
	}

	_, err = _do.Select(company.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <companies> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <companies> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <companies> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <companies> fail:", err)
	}
	
	_, err = _do.ScanByPage(&model.Company{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <companies> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <companies> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <companies> fail:", err)
	}
	
	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <companies> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <companies> fail:", err)
	}
	
	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <companies> fail:", err)
	}
}
