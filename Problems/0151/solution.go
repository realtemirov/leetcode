package problem0151

func ReverseWords(s string) string {
	res := make([]byte, 0, len(s))
	word := []byte{}
	hasWord := false
	count := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			hasWord = true
			word = append(word, byte(s[i]))
			count++
		} else {
			if hasWord {
				res = append(res, byte(' '))

				for count != 0 {
					res = append(res, word[count-1])
					count--
				}
				word = []byte{}
			}
			hasWord = false
		}
	}

	if count != 0 {
		res = append(res, byte(' '))
		for count != 0 {
			res = append(res, word[count-1])
			count--
		}
	}

	if res[0] == ' ' {
		res = res[1:]
	}
	return string(res)
}
