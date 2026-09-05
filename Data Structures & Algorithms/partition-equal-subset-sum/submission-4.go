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
		if dp[target]==true {
			return true
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
/*
1,2,3,4,5,7
target = 22/2 = 11
dp[0] = 1

num = 1
loop 11-1
dp[11] = dp[11-1] = false
dp[10] = dp[10-1] = false
...
dp[1] = dp[1-1] = true

num = 2
dp[11] = dp[11-2] = false
...
dp[3] = dp[3-2] = true
dp[2] = dp[2-2] = true

num = 3
dp[6] = dp[6-3] = true
dp[5] = dp[5-3] = true

num = 4
dp[10] = dp[10-4] = true

*/