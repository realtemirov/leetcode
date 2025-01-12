package problem0485

func FindMaxConsecutiveOnes(nums []int) int {
	maxNumber := 0
	count := 0
	for _, num := range nums {
		if num == 1 {
			count++
		} else {
			if count > maxNumber {
				maxNumber = count
			}
			count = 0
		}
	}

	if count > maxNumber {
		maxNumber = count
	}

	return maxNumber
}
