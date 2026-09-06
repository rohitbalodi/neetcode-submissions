func jump(nums []int) int {
	if len(nums)==1 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = 0
	for i, num := range nums {
		for j:=i+1; j<i+1+num && j<len(dp); j++ {
			if dp[j]==0 {
				dp[j] = dp[i]+1
			} else {
				dp[j] = min(dp[j], dp[i]+1)
			}
		}
	}

	return dp[len(dp)-1]
}
