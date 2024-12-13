package problem1768

import "strings"

func MergeAlternately(word1 string, word2 string) string {
	str := strings.Builder{}
	n := max(len(word1), len(word2))

	for i := 0; i < n; i++ {
		if i < len(word1) {
			str.WriteByte(word1[i])
		}

		if i < len(word2) {
			str.WriteByte(word2[i])
		}
	}

	return str.String()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
