package problem0350

func Intersect(nums1 []int, nums2 []int) []int {
	var m map[int]int
	if len(nums1) < len(nums2) {
		m = make(map[int]int, len(nums1))
		for _, num := range nums1 {
			m[num]++
		}
	} else {
		m = make(map[int]int, len(nums2))
		for _, num := range nums2 {
			m[num]++
		}
	}

	if len(nums1) < len(nums2) {
		nums1 = nums2
	}

	nums := make([]int, 0, len(m))
	for _, num := range nums1 {
		if m[num] != 0 {
			nums = append(nums, num)
			m[num]--
		}
	}
	return nums
}
