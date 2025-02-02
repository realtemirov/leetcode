package problem1752

import "sort"

func Check2(nums []int) bool {
	arr := make([]int, len(nums))
	copy(arr, nums)
	sort.Ints(arr)

	for i := 0; i < len(nums); i++ {
		if compare(arr, append(nums[i:], nums[:i]...)) {
			return true
		}
	}

	return false
}

func compare(arr, nums []int) bool {
	if len(arr) != len(nums) {
		return false
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != nums[i] {
			return false
		}
	}
	return true
}

func Check1(nums []int) bool {
	count := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] > nums[(i+1)%n] {
			count++
		}
		if count > 1 {
			return false
		}
	}
	return true
}
