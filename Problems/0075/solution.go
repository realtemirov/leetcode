package problem0075

func SortColors(nums []int) {
	zero, one, two := 0, 0, 0
	for _, v := range nums {
		if v == 0 {
			zero++
		} else if v == 1 {
			one++
		} else {
			two++
		}
	}

	for i := range nums {
		if i < zero {
			nums[i] = 0
		} else if i < zero+one {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
}
