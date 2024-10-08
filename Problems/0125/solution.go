package problem0125

import (
	"fmt"
	"strings"
)

func IsPalindrome(s string) bool {
		w := make([]string, 0, len(s))

		for _, v := range s {
			if ("a" <= strings.ToLower(string(v)) && strings.ToLower(string(v)) <= "z") ||
				("0" <= strings.ToLower(string(v)) && strings.ToLower(string(v)) <= "9") {
				w = append(w, strings.ToLower(string(v)))
			}
		}
		fmt.Println(w)
		for i := 0; i < len(w); i++ {
			if w[i] != w[len(w)-1-i] {
				return false
			}
		}

		return true
	}
