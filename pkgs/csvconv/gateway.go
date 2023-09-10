package csvconv

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
