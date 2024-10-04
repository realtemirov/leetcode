package problem0014

import "strings"

func LongestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
		}
	}

	return prefix
}
