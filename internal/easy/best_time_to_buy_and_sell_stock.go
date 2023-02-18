package easy

func MaxProfit(prices []int) int {
	min := prices[0]
	max := 0
	for i := 1; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		}
		if prices[i]-min > max {
			max = prices[i] - min
		}
	}

	return max
}
