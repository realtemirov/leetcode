package problem0189

import "fmt"

func Rotate(nums []int, k int) {
	n := len(nums)
	k %= n

	reverse(0, n-1, nums)
	reverse(0, k-1, nums)
	reverse(k, n-1, nums)
}

func reverse(start, end int, nums []int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
	fmt.Println(nums)
}
