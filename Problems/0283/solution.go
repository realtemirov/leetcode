package problem0283

func MoveZeroes(nums []int) []int {
	idx := 0
	for _, num := range nums {
		if num != 0 {
			nums[idx] = num
			idx++
		}
	}

	for i := idx; i < len(nums); i++ {
		nums[i] = 0
	}

	return nums
}
