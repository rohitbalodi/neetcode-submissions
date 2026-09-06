func maxSubArray(nums []int) int {
	currSum, best := -10001, -10001

	for _, num := range nums {
		currSum = max(currSum+num, num)
		best = max(best, currSum)
	}

	return best
    
}
