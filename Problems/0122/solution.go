package problem0122

func MaxProfit(prices []int) int {
	profit := 0
	min := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			profit += prices[i] - min
			min = prices[i]
		}
	}
	return profit
}
