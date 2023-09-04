package main

import (
	"io"
	"os"
	"path/filepath"
	"sample/go-gorm-example/utgen/tmpl"
	"text/template"
	"unicode"

	"github.com/gertd/go-pluralize"
)

func main() {
	gen([]string{"user", "company"}, "model", "./infra/dao")
}

func capitalize(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func filewriter(fname string) io.Writer {
	f, err := os.Create(fname)
	if err != nil {
		panic(err)
	}
	return f
}

func joinpath(base, fname string) string {
	return filepath.Join(base, fname)
}

func gen(tables []string, pkgname string, output string) {
	plu := pluralize.NewClient()

	query := map[string]interface{}{}
	querytable := []map[string]string{}
	for _, table := range tables {
		querytable = append(querytable, map[string]string{"ModelStructName": capitalize(table)})
	}
	query["Data"] = querytable
	var err error
	err = render(tmpl.QueryMethodTest, filewriter(joinpath(output, "gen_test.go")), query)
	if err != nil {
		panic(err)
	}

	for _, table := range tables {
		data := map[string]interface{}{
			"TableName":       plu.Plural(table),
			"QueryStructName": table,
			"Package":         pkgname,
			"ModelStructName": capitalize(table),
		}
		err = render(tmpl.CRUDMethodTest, filewriter(joinpath(output, plu.Plural(table)+".gen_test.go")), data)
		if err != nil {
			panic(err)
		}
	}
}

func render(tmpl string, wr io.Writer, data interface{}) error {
	t, err := template.New(tmpl).Parse(tmpl)
	if err != nil {
		return err
	}
	return t.Execute(wr, data)
}
