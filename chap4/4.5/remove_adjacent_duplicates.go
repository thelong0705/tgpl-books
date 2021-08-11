package remove_adjacent_duplicates

func RemoveAdjacentDuplicates(input []string) []string {
	output := make([]string, 0)

	for i, str := range input {
		if i+1 < len(input) && str == input[i+1] {
			continue
		}

		output = append(output, str)
	}
	return output
}
