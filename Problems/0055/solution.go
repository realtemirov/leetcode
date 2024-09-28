package problem0055

func CanJump(nums []int) bool {
	jump := 0
	for i, num := range nums {
		if i > jump {
			return false
		}

		if i+num > jump {
			jump = i + num
		}
	}

	return true
}
