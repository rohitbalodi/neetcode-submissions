func jump(nums []int) int {
	if len(nums)==1 {
		return 0
	}

	currentLevel, farthest, jumps := 0,0,0
    // last index does not require jumps
    for i:=0; i<len(nums)-1; i++ {
        farthest = max(farthest, i+nums[i])
        if i==currentLevel {
            jumps++
            currentLevel = farthest
        }
    }

    return jumps
}
