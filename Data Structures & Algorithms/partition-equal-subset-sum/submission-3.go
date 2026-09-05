func canPartition(nums []int) bool {
	// if the sum of all elements is odd,
	// then its not possible to divide it into
	// 2 equal halves
	total := sum(nums)
	if total%2==1 {
		return false
	}

	// This is the target for each half to achieve
	target := total/2
	dp := make([]bool, target+1)
	dp[0] = true

	for _, num := range nums {
		for sum:= target; sum>=num; sum-- {
			dp[sum] = dp[sum] || dp[sum-num]
		}
	}

	return dp[target]
}

func sum(nums []int) int {
	total := 0
	for _,v := range(nums) {
		total+=v
	}

	return total
}
