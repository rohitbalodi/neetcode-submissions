func subsetsWithDup(nums []int) [][]int {
	output := [][]int{}
	sort.Ints(nums)

	var backtrack func(index int, curr []int)
	backtrack = func(index int, curr []int) {
		temp := append([]int{}, curr...)
		output = append(output, temp)
		
		for i:=index; i<len(nums); i++ {
			if i>index && nums[i]==nums[i-1] {
				continue
			}
			curr = append(curr, nums[i])
			backtrack(i+1, curr)
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0, []int{})
	return output

}
