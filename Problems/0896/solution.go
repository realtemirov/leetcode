package problem0896

func IsMonotonic(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	m := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			if m == 0 {
				m = 1
			} else if m == -1 {
				return false
			}
		} else if nums[i] < nums[i-1] {
			if m == 0 {
				m = -1
			} else if m == 1 {
				return false
			}
		}
	}

	return true
}
