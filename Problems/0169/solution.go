package problem0169

func MajorityElement(nums []int) int {
	var majority, count int

	for _, num := range nums {
		if count == 0 {
			majority = num
		}
		if num == majority {
			count++
		} else {
			count--
		}
	}

	return majority
}
