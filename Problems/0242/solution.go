package problem0242

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charSpace := make(map[rune]int, len(s))

	for _, r := range s {
		charSpace[r] = charSpace[r] + 1
	}

	for _, r := range t {
		if charSpace[r] <= 0 {
			return false
		}

		charSpace[r] = charSpace[r] - 1
	}

	return true
}
