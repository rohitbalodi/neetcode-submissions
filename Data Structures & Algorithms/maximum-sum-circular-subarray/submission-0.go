func maxSubarraySumCircular(nums []int) int {
	maxSum, minSum, best, worst, total := nums[0], nums[0], nums[0], nums[0], nums[0]

	// find the best sum in a straight array
	// find the worst sum in a straight array
	// subtract worst sum from the total to
	// get best circular array sum
    for i:=1; i<len(nums); i++ {
        total+=nums[i]
        maxSum = max(nums[i], maxSum+nums[i])
        best = max(best, maxSum)

        minSum = min(nums[i], minSum+nums[i])
        worst = min(worst, minSum)
    }

    // incase all numbers are negative, this is the best you got
    if best < 0 {
        return best
    }

    // total - worst removes the worst part from the circular array
    // and gives the best circular array sum
    return max(best, total - worst)
	
}
