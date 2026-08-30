func maxProfit(prices []int) int {
	maxSoFar := prices[len(prices)-1]

	best := 0
	for i:= len(prices)-1; i>=0; i-- {
		if prices[i] > maxSoFar {
			maxSoFar = prices[i]
			prices[i] = 0
		} else {
			prices[i] = maxSoFar - prices[i]
			best = max(best, prices[i])
		}
	}
	
	return best
}
