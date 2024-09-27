package problem0080

func RemoveDuplicates(nums []int) int {
	idx := 0
	exist := true
	for i := 1; i < len(nums); i++ {
		if nums[idx] == nums[i] {
			if exist {
				idx++
				nums[idx] = nums[i]
				exist = false
			}
		} else {
			idx++
			nums[idx] = nums[i]
			exist = true
		}
	}

	return idx + 1
}
