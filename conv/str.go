package conv

func StrP(s string) *string {
	return &s
}

func Str(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
