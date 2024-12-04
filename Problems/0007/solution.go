package problem0007

import "math"
func Reverse(x int) int {
    isNegative := false

    if x < 0 {
        isNegative = true
        x *= -1
    }

    num := 0

    for x != 0 {
        num *= 10
        num += x%10
        x /= 10
    }

    if isNegative {
        num *= -1
    }

    if math.MinInt32 <= num && num <= math.MaxInt32 {
		return num
	}
    
    return 0
}