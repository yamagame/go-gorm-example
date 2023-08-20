package model

import (
	"fmt"
	"sample/go-gorm-example/conv"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      *string
	NameKana  *string
	CompanyID *uint
	Company   Company
}

func (x User) String() string {
	return fmt.Sprintf("{ID: %d, Name: %s, CompanyID: %d, Company: %s}", x.ID, conv.Str(x.Name), conv.Uint(x.CompanyID), x.Company)
}
