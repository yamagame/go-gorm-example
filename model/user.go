package model

import (
	"fmt"
	"sample/go-gorm-example/conv"

	"gorm.io/gorm"
)

type Role int32

type User struct {
	gorm.Model
	Name      *string
	NameKana  *string
	Age       *uint
	Role      *Role
	CompanyID *uint
	Company   *Company
}

func (x User) String() string {
	return fmt.Sprintf("{ID: %d, Name: %s, Age: %d, CompanyID: %d, Company: %s}", x.ID, conv.Str(x.Name), conv.Uint(x.Age), conv.Uint(x.CompanyID), x.Company)
}
