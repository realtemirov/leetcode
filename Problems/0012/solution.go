package problem0012

import "strings"

func IntToRoman(num int) string {
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := strings.Builder{}

	i := 0
	for num > 0 {
		for num >= nums[i] {
			num -= nums[i]
			roman.WriteString(romans[i])
		}
		i++
	}

	return roman.String()
}
