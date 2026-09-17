func findTargetSumWays(nums []int, target int) int {
	myMap := map[int]int {0:1}
	var backtrack func(index int, dp map[int]int) map[int]int
	backtrack = func(index int, dp map[int]int) map[int]int{
		if index==len(nums) {
			return dp
		}

		tempMap := map[int]int{}
		for key, value := range(dp) {
			// add and subtract current number, also add in map
			tempMap[key+nums[index]]+=value
			tempMap[key-nums[index]]+=value
		}
		
		return backtrack(index+1, tempMap)
	}

	final := backtrack(0, myMap)
	return final[target]
}

