func lengthOfLIS(nums []int) int {
	table := make([]int, len(nums))
	for i:=0; i<len(table); i++ {
		table[i] = 1
	}

	for i:=1; i<len(nums); i++ {
		curr := nums[i]
		for j:=i-1; j>=0; j-- {
			if curr > nums[j] {
				if table[j] >= table[i] {
					table[i] = table[j]+1
				}
			}
		}
	}

	fmt.Println(table)

	maxE := 0
	for i:=0; i<len(table); i++ {
		maxE = max(table[i], maxE)
	}
    
	return maxE
}
