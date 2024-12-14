package problem0066

func PlusOne(digits []int) []int {
	n := len(digits) - 1

	for i := n; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
			}
		} else {
			digits[i]++
			break
		}
	}

	return digits
}
