func canJump(nums []int) bool {
	// instead of maintaining an array dp we can
	// just use a variable to track how farthest we can jump

	if len(nums)==1 {
		return true
	}

	farthest := 0
	for i:=0; i<len(nums); i++ {
		if i>farthest {
			return false
		}

		farthest = max(farthest, i+nums[i])
		if farthest >=len(nums)-1 {
			return true
		}
	}

	return false
}
