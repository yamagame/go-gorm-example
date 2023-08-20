package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CRUDInterface[T any] interface {
	Create(record *T) (*T, error)
	Update(record *T) (*T, error)
	Find(record *T) (*T, error)
	List(record *T) ([]*T, error)
	Delete(record *T) error
}

type CRUDRepository[T any] struct {
	DB *gorm.DB
}

func (x *CRUDRepository[T]) Create(record *T) (*T, error) {
	err := x.DB.Create(record).Error
	return record, err
}

func (x *CRUDRepository[T]) Update(record *T) (*T, error) {
	err := x.DB.Clauses(clause.Returning{}).Save(record).Error
	return record, err
}

func (x *CRUDRepository[T]) Find(record *T) (*T, error) {
	var val T
	err := x.DB.Preload(clause.Associations).Where(record).Take(&val).Error
	return &val, err
}

func (x *CRUDRepository[T]) List(record *T) ([]*T, error) {
	var records []*T
	err := x.DB.Preload(clause.Associations).Where(record).Find(&records).Error
	return records, err
}

func (x *CRUDRepository[T]) Delete(record *T) error {
	return x.DB.Where(record).Delete(record).Error
}
