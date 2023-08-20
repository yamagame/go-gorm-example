package main

import (
	"errors"
	"fmt"
	"sample/go-gorm-example/conv"
	"sample/go-gorm-example/infra"
	"sample/go-gorm-example/model"
	"sample/go-gorm-example/repository/company"
	"sample/go-gorm-example/repository/user"

	"gorm.io/gorm"
)

func main() {
	db := infra.DB()

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Company{})

	db.Transaction(func(tx *gorm.DB) error {
		companyRepository := company.NewRepository(tx)
		userRepository := user.NewRepository(tx)

		// 会社Aの作成
		companyA, err := companyRepository.Create(&model.Company{
			Code: conv.StrP("001"),
			Name: conv.StrP("会社A"),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(companyA)

		// 会社Aの社員Aの作成
		user, err := userRepository.Create(&model.User{
			Name:    conv.StrP("社員A"),
			Company: *companyA,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(user)

		return errors.New("rollback")
	})
}
