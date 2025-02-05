package problem1800

func MaxAscendingSum(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}
	sum1, sum2 := nums[0], 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			sum1 += nums[i]
		} else {
			sum2 = max(sum1, sum2)
			sum1 = 0
			sum1 += nums[i]
		}
	}

	sum2 = max(sum1, sum2)
	return sum2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
