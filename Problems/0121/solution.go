package problem0121

func MaxProfit(prices []int) int {
	max, min := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		profit := prices[i] - min
		if prices[i] < min {
			min = prices[i]
		} else if profit > max {
			max = profit
		}
	}

	return max
}
