package problem0015

func ThreeSum(nums []int) [][]int {
	quickSort(nums, 0, len(nums)-1)

	resp := [][]int{}

	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := num + nums[l] + nums[r]
			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				resp = append(resp, []int{num, nums[l], nums[r]})
				l++

				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}
	}

	return resp
}

func quickSort(nums []int, left, right int) {
	if left < right {
		p := partition(nums, left, right)
		quickSort(nums, left, p-1)
		quickSort(nums, p+1, right)
	}
}

func partition(nums []int, left, right int) int {
	i := left - 1
	p := nums[right]

	for j := left; j < right; j++ {
		if nums[j] <= p {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	i++
	nums[right], nums[i] = nums[i], nums[right]

	return i
}
