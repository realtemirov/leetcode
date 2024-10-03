package problem0134

func CanCompleteCircuit(gas []int, cost []int) int {
	start, fill, total := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		total += gas[i] - cost[i]
		fill += gas[i] - cost[i]
		if fill < 0 {
			fill = 0
			start = i + 1
		}
	}
	if total < 0 {
		return -1
	}
	return start % len(gas)
}
