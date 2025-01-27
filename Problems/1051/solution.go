package problem1051

import "sort"

func HeightChecker(heights []int) int {
	nums := make([]int, len(heights))
	copy(nums, heights)

	sort.Ints(nums)
	count := 0

	for i, num := range nums {
		if num != heights[i] {
			count++
		}
	}

	return count
}
