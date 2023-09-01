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
	// UserをAutoMigrateすれば自動的にCompanyができるため不要
	// db.AutoMigrate(&model.Company{})

	db.Transaction(func(tx *gorm.DB) error {
		companyRepo := company.NewRepository(tx)
		userRepo := user.NewRepository(tx)

		var err error

		// 会社Aの作成
		companyA, err := companyRepo.Create(&model.Company{
			Code: conv.StrP("001"),
			Name: conv.StrP("会社A"),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(companyA)

		// 会社Bの作成
		companyB, err := companyRepo.Create(&model.Company{
			Code: conv.StrP("002"),
			Name: conv.StrP("会社B"),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(companyB)

		names := []string{
			"社員A",
			"社員B",
			"社員C",
		}

		// 会社Aの社員の作成
		for _, name := range names {
			user, err := userRepo.Create(&model.User{
				Name:    conv.StrP(name + "-A"),
				Company: companyA,
			})
			if err != nil {
				panic(err)
			}
			fmt.Println(user)
		}

		// 会社Bの社員の作成
		for _, name := range names {
			user, err := userRepo.Create(&model.User{
				Name:    conv.StrP(name + "-B"),
				Company: companyB,
			})
			if err != nil {
				panic(err)
			}
			fmt.Println(user)
		}

		// 社員A-Aを検索
		userA1, err := userRepo.Find(&model.User{Name: conv.StrP("社員A-A")})
		if err != nil {
			panic(err)
		}

		// 社員A-Aの名前と年齢変更
		userA1.Name = conv.StrP("社員X-A")
		userA1.Age = conv.UintP(34)
		userA2, err := userRepo.Update(userA1)
		if err != nil {
			panic(err)
		}
		fmt.Println(userA2)

		// 会社A社員一覧
		usersA, err := userRepo.List(&model.User{CompanyID: &companyA.ID})
		if err != nil {
			panic(err)
		}
		fmt.Println(usersA)

		// 会社Bの社員Bを削除
		err = userRepo.Delete(&model.User{Name: conv.StrP("社員B-B"), CompanyID: &companyB.ID})
		if err != nil {
			panic(err)
		}

		// 会社B社員一覧
		usersB, err := userRepo.List(&model.User{CompanyID: &companyB.ID})
		if err != nil {
			panic(err)
		}
		fmt.Println(usersB)

		return errors.New("rollback")
	})
}
