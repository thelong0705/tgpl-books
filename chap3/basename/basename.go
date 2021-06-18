package basename

import "strings"

func Basename(s string) string {
	lastIndexOfSlash := strings.LastIndex(s, "/")
	s = s[lastIndexOfSlash+1:]
	lastIndexOfDot := strings.LastIndex(s, ".")

	if lastIndexOfDot >= 0 {
		s = s[:lastIndexOfDot]
	}

	return s
}
