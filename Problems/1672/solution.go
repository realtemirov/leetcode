package problem1672

func MaximumWealth(accounts [][]int) int {
	max := 0
	for _, i := range accounts {
		sum := 0
		for _, j := range i {
			sum += j
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
