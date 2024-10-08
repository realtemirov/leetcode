package problem0015

import "sort"

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			threeSum := num + nums[l] + nums[r]
			if threeSum > 0 {
				r--
			} else if threeSum < 0 {
				l++
			} else {
				res = append(res, []int{num, nums[l], nums[r]})
				l += 1

				for nums[l] == nums[l-1] && l < r {
					l += 1
				}
			}
		}
	}

	return res
}
