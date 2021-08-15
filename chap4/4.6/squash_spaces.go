package squash_spaces

import (
	"unicode"
	"unicode/utf8"
)

func SquashSpaces(input []byte) []byte {
	lastSize := 0
	sizeCounter := 0
	l := len(input)
	var lastRune rune

	for {
		r, size := utf8.DecodeRune(input[lastSize:])

		if unicode.IsSpace(r) && unicode.IsSpace(lastRune) {
			input = append(input[:lastSize], input[lastSize+size:]...)

			if lastRune != rune(32) {
				input[lastSize-utf8.RuneLen(lastRune)] = ' '
				input = append(input[:lastSize-utf8.RuneLen(lastRune)+1], input[lastSize:]...)
				lastSize = lastSize -utf8.RuneLen(lastRune)+1
				lastRune = rune(32)
			}

		} else {
			lastSize = lastSize + size
			lastRune = r
		}
		sizeCounter = sizeCounter + size
		if sizeCounter == l {
			break
		}
	}

	return input
}
