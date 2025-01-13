package problem3223

func MinimumLength(s string) int {
	if len(s) < 3 {
		return len(s)
	}

	m := [26]int{}

	for _, c := range s {
		m[c-'a']++
	}

	sum := 0
	for _, val := range m {
		if val == 0 {
			continue
		}
		if val&1 == 1 {
			sum += 1
		} else {
			sum += 2
		}
	}

	return sum
}
