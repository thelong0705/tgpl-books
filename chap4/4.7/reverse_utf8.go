package reverse_utf8

import "unicode/utf8"

func rev(s []byte) {
	l := len(s)
	for i := 0; i < len(s)/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
}

// ReverseUTF8 reverse the characters of a []byte slice that represents a UTF-8-encoded string, in place
func ReverseUTF8(s []byte) {
	i := 0

	for i < len(s) {
		_, size := utf8.DecodeRune(s[i:])
		rev(s[i : i+size])
		i = i + size
	}

	rev(s)
}
