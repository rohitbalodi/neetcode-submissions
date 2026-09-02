func subsets(nums []int) [][]int {
	output := [][]int {}

	var backtrack func(index int, curr []int)
	backtrack = func(index int, curr []int) {
		if len(curr)==len(nums) {
			temp := append([]int{}, curr...)
			output = append(output, temp)
			return
		}

		temp := append([]int{}, curr...)
		output = append(output, temp)
		for i:= index; i<len(nums); i++ {
			curr = append(curr, nums[i])
			backtrack(i+1, curr)
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0, []int{})
	return output
}
