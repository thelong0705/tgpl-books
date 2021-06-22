package anagram

func IsAnagram(s1, s2 string) bool {
	return sort(s1) == sort(s2)
}

func sort(s string) string {
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}
