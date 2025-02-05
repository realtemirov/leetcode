package problem3105

func LongestMonotonicSubarray(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	inc, dec, maxm := 1, 1, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			inc++
			dec = 1
		}
		if nums[i] < nums[i-1] {
			dec++
			inc = 1
		}
		if nums[i] == nums[i-1] {
			inc = 1
			dec = 1
		}
		maxm = max(inc, dec, maxm)
	}
	return maxm
}
func max(a, b, c int) int {
	if a > b {
		return a
	}
	if b > c {
		return b
	}
	return c
}
