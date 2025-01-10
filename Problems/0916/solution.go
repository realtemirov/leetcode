package problem0916

func WordSubsets(a []string, b []string) []string {
	bMax := count("")

	for _, word := range b {
		countOfChar := count(word)
		for i := 0; i < 26; i++ {
			bMax[i] = max(bMax[i], countOfChar[i])
		}
	}

	resp := make([]string, 0, len(a))

	for _, word := range a {
		aCount := count(word)
		if check(bMax, aCount) {
			resp = append(resp, word)
		}
	}

	return resp
}

func check(bMax, aCount []int) bool {
	for i := 0; i < 26; i++ {
		if bMax[i] > aCount[i] {
			return false
		}
	}

	return true
}

func count(s string) []int {
	chars := make([]int, 26)

	for _, char := range s {
		chars[char-'a']++
	}

	return chars
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
