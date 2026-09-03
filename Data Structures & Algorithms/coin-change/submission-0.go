func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i:=1; i<len(dp); i++ {
		dp[i] = amount+1
	}
	
	for currAmount:=1; currAmount<=amount; currAmount++ {
		for j:=0; j<len(coins); j++ {
			curr := coins[j]
			if currAmount < curr {
				continue
			}

			dp[currAmount] = min(dp[currAmount], dp[currAmount - curr] + 1)
		}
	}

	if dp[amount]>amount {
		return -1
	}

	return dp[amount]
    
}
