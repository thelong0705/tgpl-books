package comma

import (
	"bytes"
	"strings"
)

func comma(s string) string {
	for i := len(s) - 3; i > 0; i = i - 3 {
		s = s[:i] + "," + s[i:]
	}
	return s
}

func recursiveComma(s string) string {
	if len(s) <= 3 {
		return s
	}
	return recursiveComma(s[:len(s)-3]) + "," + s[len(s)-3:]
}

func commaWithByteBuffer(s string) string {
	var buf bytes.Buffer
	remainder := len(s) % 3
	buf.WriteString(s[:remainder])

	for i := remainder; i < len(s); i = i + 3 {
		if i != 0 {
			buf.WriteByte(',')

		}
		buf.WriteString(s[i : i+3])
	}

	return buf.String()
}

func forFloatComma(s string) string {
	dot := strings.LastIndex(s, ".")
	var floatPart string
	var sign string
	if dot != -1 {
		floatPart = s[dot:]
		s = s[0:dot]
	}
	if s[0] == '+' || s[0] == '-' {
		sign = s[0:1]
		s = s[1:]
	}

	return sign + commaWithByteBuffer(s) + floatPart
}
