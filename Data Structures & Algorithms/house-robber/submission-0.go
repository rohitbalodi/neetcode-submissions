func rob(nums []int) int {
	if len(nums)<3 {
		if len(nums) == 1 {
			return nums[0]
		}
		return max(nums[0], nums[1])
	}
    
	dp := make([]int, len(nums))
	dp[0], dp[1], dp[2] = nums[0], nums[1], nums[0]+nums[2]

	for i:=3; i<len(nums); i++ {
		dp[i] = max(dp[i-2], dp[i-3]) + nums[i]
	}

	maxE := 0
	for i:=0; i<len(dp); i++ {
		maxE = max(maxE, dp[i])
	}

	return maxE
}
