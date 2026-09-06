func canJump(nums []int) bool {
	// if only 1 number is present, you are already at end
	if len(nums)==1 {
		return true
	}

	// if more than 1 number and first is 0, you can't proceed
	if nums[0]==0 {
		return false
	}

	dp := make([]bool, len(nums)+1)
	// 0th index is reachable since its not zero
	dp[0] = true

	// traverse every element num and then
	// i+1, i+2, ... i+num are all reachable
	// if you encounter any point in table which is unreachable
	// then return false since there is a gap  
	for i:=0; i<len(nums); i++ {
		if dp[i]== false {
			return false
		}
		curr := nums[i]
		for j:=i+1; j<(i+1+curr) && j<len(nums); j++ {
			dp[j] = true
		}
	}

	return dp[len(nums)-1]
}
