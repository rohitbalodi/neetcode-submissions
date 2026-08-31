func combinationSum(nums []int, target int) [][]int {
	output := [][]int {}

	var backtrack func(i int, curr []int, remaining int)
	backtrack = func(i int, curr []int, remaining int) {
		if remaining ==0 {
			temp := append([]int{}, curr...)
			output = append(output, temp)
			return
		}

		if remaining < 0 || i>=len(nums) {
			return
		}

		curr = append(curr, nums[i])
		backtrack(i, curr, remaining - nums[i])
		curr = curr[:len(curr)-1]
		backtrack(i+1, curr, remaining)
	}

	backtrack(0, []int{}, target)
	return output
    
}
