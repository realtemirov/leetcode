package problem1769

func MinOperations(boxes string) []int {
	resp := make([]int, len(boxes))
	var (
		bl, ml = 0, 0
		br, mr = 0, 0
		j      int
	)
	n := len(boxes)
	for i := 0; i < n; i++ {
		resp[i] += ml
		bl += int(boxes[i] - '0')
		ml += bl

		j = n - 1 - i
		resp[j] += mr
		br += int(boxes[j] - '0')
		mr += br
	}

	return resp
}
