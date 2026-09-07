func dailyTemperatures(temperatures []int) []int {
	stack := [][]int{{101,-1}}
	output := make([]int, len(temperatures))

	for i:=0; i<len(temperatures); i++ {
		if stack[len(stack)-1][0] < temperatures[i] {
			for stack[len(stack)-1][0] < temperatures[i] {
				index := stack[len(stack)-1][1]
				output[index] = i-index
				stack = stack[:len(stack)-1]
			}
		}
		stack = append(stack, []int{temperatures[i], i})
	}

	return output
}
