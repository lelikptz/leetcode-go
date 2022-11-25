package simple

func MaxProfit(prices []int) int {
	prepared := preparePricesList(prices)

	return calculateMaxProfit(prepared, 0)
}

func preparePricesList(prices []int) []int {
	prepared := make([]int, 0, len(prices))

	for i := 0; i < len(prices); i++ {
		if i == 0 || i == len(prices)-1 || (!(prices[i] > prices[i-1] && prices[i] < prices[i+1]) && !(prices[i] <= prices[i-1] && prices[i] >= prices[i+1])) {
			prepared = append(prepared, prices[i])
		}
	}

	return prepared
}

func calculateMaxProfit(prices []int, index int) int {
	if index >= len(prices) {
		return 0
	}

	max := prices[index]
	maxIndex := index
	for i := index; i < len(prices); i++ {
		if prices[i] > max {
			max = prices[i]
			maxIndex = i
		}
	}

	min := prices[index]
	for i := index; i < maxIndex; i++ {
		if prices[i] < min {
			min = prices[i]
		}
	}

	if maxIndex == index {
		maxIndex = index + 1
	}

	x := calculateMaxProfit(prices, maxIndex)
	if max-min > x {
		return max - min
	} else {
		return x
	}
}
