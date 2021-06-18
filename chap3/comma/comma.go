package comma

func comma(s string) string {
	for i := len(s) - 3; i > 0; i = i - 3 {
		s = s[:i] + "," + s[i:]
	}
	return s
}
