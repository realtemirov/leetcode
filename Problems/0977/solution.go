package problem0977

func SortedSquares(nums []int) []int {
	l, r := 0, len(nums)-1
	resp := make([]int, 0, len(nums))

	for l <= r {
		if nums[l]*nums[l] >= nums[r]*nums[r] {
			resp = append(resp, nums[l]*nums[l])
			l++
		} else {
			resp = append(resp, nums[r]*nums[r])
			r--
		}
	}

	for i, j := 0, len(resp)-1; i < j; i, j = i+1, j-1 {
        resp[i], resp[j] = resp[j], resp[i]
    }

	return resp
}