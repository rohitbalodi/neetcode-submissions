func merge(intervals [][]int) [][]int {
	stack := [][]int{{-1,-1}}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i:=0; i<len(intervals); i++ {
		top := stack[len(stack)-1]
		if intervals[i][0]<=top[1] {
			top[1] = max(top[1],intervals[i][1])
		} else {
			stack = append(stack, intervals[i])
		}
	}

	fmt.Println(stack)
	return stack[1:]
	
    
}
