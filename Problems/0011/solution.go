package problem0011

func MaxArea(h []int) int {
	area := 0
	l, r := 0, len(h)-1

	for l <= r {
		area = max(area, (r-l)*(min(h[l], h[r])))

		if h[l] < h[r] {
			l++
		} else {
			r--
		}
	}

	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
