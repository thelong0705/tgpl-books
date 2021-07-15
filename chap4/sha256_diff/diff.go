package sha256_diff

func Diff(input1, input2 [32]byte) int {
	var diff int
	for i, _ := range input1 {
		diff += popCount(input1[i] ^ input2[i])
	}
	return diff
}

func popCount(x byte) int {
	count := 0

	for ; x > 0; {
		count += 1
		x = x & (x - 1)
	}

	return count
}
