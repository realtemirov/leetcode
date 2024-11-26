package problem0034

func SearchRange(nums []int, target int) []int {
	l := LowerBound(nums, target)
	if l == len(nums) || nums[l] != target {
		return []int{-1, -1}
	}

	r := UpperBound(nums, target)
	return []int{l, r - 1}
}

func LowerBound(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}

func UpperBound(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}
