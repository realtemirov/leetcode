package problem0459

import "strings"

func RepeatedSubstringPattern(s string) bool {
	d := s + s
	return strings.Contains(d[1:len(d)-1], s)
}
