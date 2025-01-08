package problem3042

import "strings"

func CountPrefixSuffixPairs(words []string) int {
	count := 0

	for i, word := range words {
		for j := i + 1; j < len(words); j++ {
			if len(word) <= len(words[j]) && isPrefixAndSuffix(words[j], word) {
				count++
			}
		}
	}

	return count
}

func isPrefixAndSuffix(word1, word2 string) bool {
	return strings.HasPrefix(word1, word2) && strings.HasSuffix(word1, word2)
}
