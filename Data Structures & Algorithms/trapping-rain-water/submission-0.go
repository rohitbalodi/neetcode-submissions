func trap(height []int) int {
	left := make([]int, len(height))
	right := make([]int, len(height))

	maxSoFar := 0
	for i:=1; i<len(height); i++ {
		maxSoFar = max(maxSoFar, height[i-1])
		left[i] = maxSoFar
	}

	maxSoFar = 0
	for i:=len(height)-2; i>=0; i-- {
		maxSoFar = max(maxSoFar, height[i+1])
		right[i] = maxSoFar
	}

	output := 0

	for i:=0; i<len(height); i++ {
		output = output+ max(0,(min(left[i], right[i]) - height[i]))
	}
	return output


}

/*
[0,2,0,3,1,0,1,3,2,1]
[0,0,2,2,3,3,3,3,3,3]
[3,3,3,3,3,3,3,2,1,0]
[0,0,2,2,3,3,3,2,1,1]
*/