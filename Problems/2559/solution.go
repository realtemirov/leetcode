package problem2559

func VowelStrings(words []string, queries [][]int) []int {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	ans := make([]int, len(queries))
	vowelCounts := make([]int, len(words)+1)

	for i, word := range words {
		if vowels[word[0]] && vowels[word[len(word)-1]] {
			vowelCounts[i+1]++
		}
		vowelCounts[i+1] += vowelCounts[i]
	}

	for i, query := range queries {
		ans[i] = vowelCounts[query[1]+1] - vowelCounts[query[0]]
	}

	return ans
}
