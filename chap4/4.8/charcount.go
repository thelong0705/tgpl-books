package charcount

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	CharCount("abc")
}

// CharCount count and print unique and its frequencies
func CharCount(s string) {
	fmt.Println("rune freq")
	freqMap := countRuneFreq(s)
	fmt.Println("rune\tcount")
	for k, v := range freqMap {
		fmt.Printf("'%c'\t%d\n", k, v)
	}

	fmt.Println("len freq")
	fmt.Println("len\tcount")
	for l, c := range countLenFreq(s) {
		if l == 0 {
			continue
		}
		fmt.Printf("%d\t%d\n", l, c)
	}

	catsFreq := countCategoryFreq(s)
	fmt.Println("cat\tcount")
	for cat, freq := range catsFreq {
		fmt.Printf("%s\t%d\n", cat, freq)
	}
}

func countRuneFreq(s string) map[rune]int {
	freqMap := make(map[rune]int)
	for _, r := range s {
		freqMap[r]++
	}
	return freqMap
}

func countLenFreq(s string) (res [utf8.UTFMax + 1]int) {
	for _, r := range s {
		res[utf8.RuneLen(r)]++
	}
	return res
}

func countCategoryFreq(s string) map[string]int {
	catsMap := make(map[string]int)
	for _, r := range s {
		for catName, rangeTable := range unicode.Properties {
			if unicode.In(r, rangeTable) {
				if catName == "Hex_Digit" {
					fmt.Println(r)
					fmt.Printf("%c\n", r)
				}
				catsMap[catName]++
			}
		}
	}
	return catsMap
}
