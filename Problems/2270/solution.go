package problem2270

func WaysToSplitArray(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	var (
		count   = 0
		leftSum = 0
	)
	for i := 0; i < len(nums)-1; i++ {
		leftSum += nums[i]
		if leftSum >= (sum - leftSum) {
			count++
		}
	}

	return count
}
