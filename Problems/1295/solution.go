package problem1295

import "fmt"

func FindNumbers(nums []int) int {
	sum := 0
	for _, num := range nums {
		if len(fmt.Sprintf("%d", num))%2 == 0 {
			sum++
		}
	}
	return sum
}
