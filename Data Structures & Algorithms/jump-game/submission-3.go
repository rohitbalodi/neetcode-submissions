func canJump(nums []int) bool {
    dp := make([]bool, len(nums)+1)
	if len(nums)==1 {
		return true
	}
	if nums[0]==0 {
		return false
	}

	dp[0] = true

	for i:=0; i<len(nums); i++ {
		curr := nums[i]
		j:=i+1
		for j<(i+1+curr) && j<len(nums) {
			dp[j] = true
			j++
		}
		if dp[i]== false {
			return false
		}
	}

	return dp[len(nums)-1]
}
