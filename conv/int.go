package conv

func UintP(v uint) *uint {
	return &v
}

func Uint(v *uint) uint {
	if v == nil {
		return 0
	}
	return *v
}

func Int64P(v int64) *int64 {
	return &v
}
