package csvconv

import "fmt"

type GatewayInterface[T any] interface {
	ToCSV(v interface{}) (string, error)
	FromCSV(v string) (interface{}, error)
	SetInfo(mapping []*Mapping[T], record []string, objptr *T, fields map[string]interface{})
	ColumnValue(field string) string
}

type Gateway[T any] struct {
	Self    *T
	Record  []string
	Field   map[string]interface{}
	Mapping []*Mapping[T]
}

func (x *Gateway[T]) SetInfo(mapping []*Mapping[T], record []string, objptr *T, fields map[string]interface{}) {
	x.Mapping = mapping
	x.Record = record
	x.Self = objptr
	x.Field = fields
}

func (Gateway[T]) ToCSV(v interface{}) (string, error) {
	return fmt.Sprintf("%v", v), nil
}

func (Gateway[T]) FromCSV(v string) (interface{}, error) {
	return v, nil
}

func (x *Gateway[T]) ColumnValue(column string) string {
	for i, m := range x.Mapping {
		if m.Column == column {
			return x.Record[i]
		}
	}
	return ""
}

type Mapping[T any] struct {
	Column  string
	Key     string
	Gateway GatewayInterface[T]
}

////////////////////////////////////////////////////////////////////////////////
// ユーティリティ
////////////////////////////////////////////////////////////////////////////////

type EmptyString[T any] struct {
	Gateway[T]
}

func (x *EmptyString[T]) ToCSV(v interface{}) (string, error) {
	return "", nil
}

func (x *EmptyString[T]) FromCSV(v string) (interface{}, error) {
	return "", nil
}

type StaticString[T any] struct {
	Value string
	Gateway[T]
}

func (x *StaticString[T]) ToCSV(v interface{}) (string, error) {
	return x.Value, nil
}

func (x *StaticString[T]) FromCSV(v string) (interface{}, error) {
	return "", nil
}

type ConvString[T any] struct {
	To   func(m *T) (string, error)
	From func(m string) (string, error)
	Gateway[T]
}

func (x *ConvString[T]) ToCSV(v interface{}) (string, error) {
	return x.To(x.Self)
}

func (x *ConvString[T]) FromCSV(v string) (interface{}, error) {
	return x.From(v)
}
