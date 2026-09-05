func lengthOfLIS(nums []int) int {
	table := make([]int, len(nums))
	var maxE int

	for i:=0; i<len(nums); i++ {
		table[i] = 1
		curr := nums[i]
		for j:=0; j<i; j++ {
			if curr > nums[j] {
				table[i] = max(table[i], table[j]+1)
			}
		}

		maxE = max(maxE, table[i])
	}
    
	return maxE
}
