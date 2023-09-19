package main

import (
	"flag"
	"fmt"
	"sample/go-gorm-example/infra"
	"sample/go-gorm-example/migrations"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

const DB_NAME = "go-gorm-example"

func main() {
	flag.Parse()
	args := flag.Args()

	db := infra.DB()
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	driver, err := mysql.WithInstance(sqlDB, &mysql.Config{})
	if err != nil {
		panic(err)
	}
	source, err := iofs.New(migrations.FS, ".")
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithInstance("iofs", source, DB_NAME, driver)
	if err != nil {
		panic(err)
	}

	switch args[0] {
	case "up":
		if err := m.Steps(1); err != nil {
			panic(err)
		}
	case "down":
		if err := m.Steps(-1); err != nil {
			panic(err)
		}
	case "drop":
		if err := m.Drop(); err != nil {
			panic(err)
		}
	case "top":
		if err := m.Up(); err != nil {
			panic(err)
		}
	case "bottom":
		if err := m.Down(); err != nil {
			panic(err)
		}
	case "force":
		if len(args) > 1 {
			ver, err := strconv.Atoi(args[1])
			if err != nil {
				panic(err)
			}
			m.Force(ver)
		}
	case "migrate":
		if len(args) > 1 {
			ver, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				panic(err)
			}
			m.Migrate(uint(ver))
		}
	case "version":
		ver, dirty, err := m.Version()
		if err != nil {
			panic(err)
		}
		fmt.Println("version:", ver, "dirty:", dirty)
	}
}
