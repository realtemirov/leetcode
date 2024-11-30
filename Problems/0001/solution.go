package problem0001

func TwoSum(nums []int, target int) []int {
	m := map[int]int{}

	for i, num := range nums {
		if idx, ok := m[target-num]; ok {
			return []int{idx, i}
		}
		m[num] = i
	}

	return []int{}
}
