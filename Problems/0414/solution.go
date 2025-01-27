package problem0414

import "sort"

func ThirdMax(nums []int) int {
	sort.Ints(nums)
	if len(nums) < 3 {
		return nums[len(nums)-1]
	}
	cnt := 2
	i := len(nums) - 1
	for ; i > 0 && cnt != 0; i-- {

		if nums[i] != nums[i-1] {
			cnt--
		}
	}
	if cnt != 0 {
		return nums[len(nums)-1]
	}

	return nums[i]
}
