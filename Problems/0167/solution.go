package problem0167

import "fmt"

func TwoSum(numbers []int, target int) []int {
	m := map[int]int{}

	for i, v := range numbers {
		if idx, ok := m[target-v]; ok {
			fmt.Println(v, target-v, target)
			return []int{idx + 1, i + 1}
		}

		m[v] = i
	}

	return []int{}
}

func TwoSumBinarySearch(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	sum := numbers[l] + numbers[r]

	for sum != target {
		if sum > target {
			r--
		} else {
			l++
		}
		sum = numbers[l] + numbers[r]
	}
	return []int{l + 1, r + 1}
}
