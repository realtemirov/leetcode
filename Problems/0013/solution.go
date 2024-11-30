package problem0013

func RomanToInt(s string) int {
	romans := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	number := romans[s[0]]

	for i := 1; i < len(s); i++ {
		if romans[s[i-1]] < romans[s[i]] {
			number -= romans[s[i-1]]
			number += (romans[s[i]] - romans[s[i-1]])
		} else {
			number += romans[s[i]]
		}
	}

	return number
}
