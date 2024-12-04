package problem0009

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	num := 0
	k := x
	for k != 0 {
		num *= 10
		num += k % 10
		k /= 10
	}

	return x == num
}
