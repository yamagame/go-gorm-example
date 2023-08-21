package conv

func UbyteP(v uint8) *uint8 {
	return &v
}

func Ubyte(v *uint8) uint8 {
	if v == nil {
		return 0
	}
	return *v
}
