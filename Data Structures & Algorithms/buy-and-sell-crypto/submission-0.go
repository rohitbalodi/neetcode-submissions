func maxProfit(prices []int) int {
	highest := make([]int, len(prices))
	maxSoFar := prices[len(prices)-1]

	for i:= len(prices)-1; i>=0; i-- {
		if prices[i] > maxSoFar {
			maxSoFar = prices[i]
		}
		highest[i] = maxSoFar
	}

	best := 0
	for index, value := range prices {
		best = max(best, highest[index] - value)
	}
	
	return best
}
