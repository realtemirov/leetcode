package problem0274

import (
	"sort"
)

func HIndex(citations []int) int {
	var h int

	sort.Ints(citations)

	for i, v := range citations {
		h = len(citations) - i
		if v >= h {
			return h
		}
	}

	return 0
}
