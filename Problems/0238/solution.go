package problem0238

func ProductExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	l := 1
	for i, num := range nums {
		res[i] = l
		l *= num
	}

	r := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= r
		r *= nums[i]
	}

	return res
}
