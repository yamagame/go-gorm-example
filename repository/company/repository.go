package company

import (
	"sample/go-gorm-example/model"
	"sample/go-gorm-example/repository"

	"gorm.io/gorm"
)

type Repository struct {
	repository.CRUDRepository[model.Company]
}

func NewRepository(db *gorm.DB) repository.CRUDInterface[model.Company] {
	return &Repository{
		CRUDRepository: repository.CRUDRepository[model.Company]{
			DB: db,
		},
	}
}
