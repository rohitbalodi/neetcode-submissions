func permute(nums []int) [][]int {
	output := [][]int{}
	used := make([]bool, len(nums))

	var backtrack func(used []bool, curr []int)
	backtrack = func(used []bool, curr []int) {
		if len(curr)==len(nums) {
			temp := append([]int{}, curr...)
			output = append(output, temp)
			return
		}

		for i:=0; i<len(nums); i++ {
			if used[i]==true {
				continue
			}

			used[i] = true
			curr = append(curr, nums[i])
			backtrack(used, curr)
			curr = curr[:len(curr)-1]
			used[i] = false
		}
	}

	backtrack(used, []int{})
	return output

}
