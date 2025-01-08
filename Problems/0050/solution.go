package problem0050

import "math"

func MyPow(x float64, n int) float64 {
	minus := n < 0
	if minus {
		n = -n
	}
	var a float64 = 1
	for n != 0 {
		if n&1 != 0 {
			a *= x
		}
		x *= x
		n = n >> 1
	}
	if minus {
		a = 1 / a
	}

	multiplier := math.Pow(10, float64(4))
	return math.Round(a*multiplier) / multiplier
}
