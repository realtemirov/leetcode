package problem0026

func RemoveDuplicates(nums []int) int {
	idx := 1

	for i := 1; i < len(nums); i++ {
		if nums[idx-1] != nums[i] {
			nums[idx] = nums[i]
			idx++
		}
	}

	return idx
}
