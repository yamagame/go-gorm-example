package model

import (
	"fmt"
	"sample/go-gorm-example/conv"

	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Code *string
	Name *string
}

func (x Company) String() string {
	return fmt.Sprintf("{ID: %d, Code: %s, Name: %s}", x.ID, conv.Str(x.Code), conv.Str(x.Name))
}
