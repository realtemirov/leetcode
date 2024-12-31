package problem0860

func LemonadeChange(bills []int) bool {
	m := map[int]int{}

	for _, v := range bills {
		m[v]++
		if v == 10 {
			if m[5] > 0 {
				m[5]--
			} else {
				return false
			}
		} else if v == 20 {
			if m[10] >= 1 && m[5] >= 1 {
				m[10]--
				m[5]--
			} else if m[5] >= 3 {
				m[5] -= 3
			} else {
				return false
			}
		}
	}

	return true
}
