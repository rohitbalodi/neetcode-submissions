func rob(nums []int) int {
	if len(nums)<3 {
		if len(nums)==1 {
			return nums[0]
		}
		return max(nums[0], nums[1])
	}

	case1 := rob1(nums[0:len(nums)-1])
	case2 := rob1(nums[1:])

	return max(case1, case2)
    
}

func rob1(nums []int) int {
	if len(nums)<3 {
		if len(nums)==1 {
			return nums[0]
		}
		return max(nums[0], nums[1])
	}

	dp := make([]int, len(nums))
	dp[0], dp[1], dp[2] = nums[0], nums[1], nums[0] + nums[2]

	for i:=3; i<len(nums); i++ {
		dp[i] = max(dp[i-3], dp[i-2]) + nums[i]
	}

	maxE := 0
	for i:=0; i<len(dp); i++ {
		maxE = max(dp[i], maxE)
	}

	return maxE
}
