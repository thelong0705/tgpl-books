package rotate

func Rotate(s []int, n int) []int {
	return append(s[n:], s[:n]...)
}
