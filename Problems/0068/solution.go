package problem0068

import (
	"fmt"
	"strings"
)

func FullJustify(words []string, maxWidth int) []string {
	res := []string{}
	tmpLength := 0
	tmpWords := []string{}

	for _, v := range words {
		if tmpLength+len(v) >= maxWidth {
			fmt.Println("in conditino", tmpLength, tmpWords)
			res = append(res, parse(tmpWords, maxWidth))
			tmpLength = 0
			tmpWords = []string{}
		}

		tmpLength += len(v) + 1
		tmpWords = append(tmpWords, v)
		fmt.Println(v, tmpLength)
	}
	res = append(res, parse(tmpWords, maxWidth))

	return res
}
"example of text"
func parse(words []string, maxWidth int) string {
	length := 0

	for _, v := range words {
		length += len(v)
	}

	spaceSize := (maxWidth - length) / (len(words) - 1)

	res := strings.Builder{}

	for i, v := range words {
		res.WriteString(v)
		if i != len(words) {
			res.WriteString(strings.Repeat(" ", spaceSize))
		}

	}
	fmt.Println(words, maxWidth, length, spaceSize, res.String())
	return res.String()
}
