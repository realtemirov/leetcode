package problem0045

func Jump(nums []int) int {
	count := 0
	n := len(nums)
	coverage, lastJumpIdx := 0, 0

	if n == 1 {
		return 0
	}

	for i, num := range nums {
		coverage = max(coverage, i+num)

		if i == lastJumpIdx {
			lastJumpIdx = coverage
			count++

			if coverage >= n-1 {
				return count
			}
		}
	}

	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
