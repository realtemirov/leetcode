package problem0389

func FindTheDifference(s string, t string) byte {
	n := min(len(s), len(t))

	m := map[byte]int{}
	for i := 0; i < n; i++ {
		m[s[i]]++
		m[t[i]]--
	}
	m[t[len(t)-1]]--

	for k, v := range m {
		if v != 0 {
			return k
		}
	}
	return t[len(t)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
