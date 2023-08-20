package user

import (
	"sample/go-gorm-example/model"
	"sample/go-gorm-example/repository"

	"gorm.io/gorm"
)

type Repository struct {
	repository.CRUDRepository[model.User]
}

func NewRepository(db *gorm.DB) repository.CRUDInterface[model.User] {
	return &Repository{
		CRUDRepository: repository.CRUDRepository[model.User]{
			DB: db,
		},
	}
}
